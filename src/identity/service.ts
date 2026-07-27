import { randomUUID } from "node:crypto";
import type { Database } from "../db.js";
import type { Config } from "../config.js";
import { hashPassword, newOpaqueToken, tokenHash, verifyPassword } from "./crypto.js";

const SESSION_MS = 7 * 24 * 60 * 60 * 1000;
const MAX_FAILURES = 8;
const LOCK_MS = 15 * 60 * 1000;

export class IdentityService {
  constructor(private readonly db: Database, private readonly config: Config) {}

  async register(email: string, username: string, password: string, meta: RequestMeta) {
    const userId = randomUUID();
    const passwordHash = await hashPassword(password, this.config.PASSWORD_PEPPER);
    try {
      return await this.db.begin(async sql => {
        const [user] = await sql`
          INSERT INTO users (id, email, username, password_hash)
          VALUES (${userId}, ${email}, ${username}, ${passwordHash})
          RETURNING id, email::text, username::text, created_at
        `;
        await sql`INSERT INTO audit_events ${sql({
          subjectUserId: userId, eventType: "identity.registered",
          ipAddress: meta.ip, userAgent: meta.userAgent
        })}`;
        return user;
      });
    } catch (error: unknown) {
      if (isUniqueViolation(error)) throw new IdentityError("ACCOUNT_EXISTS", 409);
      throw error;
    }
  }

  async login(login: string, password: string, meta: RequestMeta) {
    const [user] = await this.db`
      SELECT id, email::text, username::text, password_hash, failed_login_count,
             locked_until, deleted_at
      FROM users WHERE email = ${login} OR username = ${login}
    `;
    const now = new Date();
    if (!user || user.deletedAt || (user.lockedUntil && user.lockedUntil > now)) {
      throw new IdentityError("INVALID_CREDENTIALS", 401);
    }
    const valid = await verifyPassword(user.passwordHash, password, this.config.PASSWORD_PEPPER);
    if (!valid) {
      await this.db`
        UPDATE users SET
          failed_login_count = failed_login_count + 1,
          locked_until = CASE WHEN failed_login_count + 1 >= ${MAX_FAILURES}
            THEN now() + ${LOCK_MS + " milliseconds"}::interval ELSE locked_until END,
          updated_at = now()
        WHERE id = ${user.id}
      `;
      await this.audit("identity.login_failed", user.id, null, meta);
      throw new IdentityError("INVALID_CREDENTIALS", 401);
    }
    return this.db.begin(async sql => {
      await sql`UPDATE users SET failed_login_count = 0, locked_until = NULL, updated_at = now() WHERE id = ${user.id}`;
      const rawToken = newOpaqueToken();
      const csrfToken = newOpaqueToken();
      const sessionId = randomUUID();
      const expiresAt = new Date(Date.now() + SESSION_MS);
      await sql`
        INSERT INTO sessions (id, user_id, token_hash, csrf_hash, user_agent, ip_address, expires_at)
        VALUES (${sessionId}, ${user.id}, ${tokenHash(rawToken, this.config.SESSION_HASH_KEY)},
          ${tokenHash(csrfToken, this.config.SESSION_HASH_KEY)}, ${meta.userAgent}, ${meta.ip}, ${expiresAt})
      `;
      await sql`INSERT INTO audit_events ${sql({
        actorUserId: user.id, subjectUserId: user.id, sessionId,
        eventType: "identity.login_succeeded", ipAddress: meta.ip, userAgent: meta.userAgent
      })}`;
      return { user: publicUser(user), sessionId, rawToken, csrfToken, expiresAt };
    });
  }

  async authenticate(rawToken: string) {
    const [row] = await this.db`
      SELECT s.id AS session_id, s.user_id, s.csrf_hash, s.expires_at,
             u.email::text, u.username::text
      FROM sessions s JOIN users u ON u.id = s.user_id
      WHERE s.token_hash = ${tokenHash(rawToken, this.config.SESSION_HASH_KEY)}
        AND s.revoked_at IS NULL AND s.expires_at > now() AND u.deleted_at IS NULL
    `;
    return row || null;
  }

  async logout(sessionId: string, meta: RequestMeta) {
    await this.db`UPDATE sessions SET revoked_at = now(), revoke_reason = 'logout' WHERE id = ${sessionId} AND revoked_at IS NULL`;
    await this.audit("identity.logout", null, sessionId, meta);
  }

  verifyCsrf(stored: Buffer, supplied: string): boolean {
    return stored.equals(tokenHash(supplied, this.config.SESSION_HASH_KEY));
  }

  private async audit(eventType: string, userId: string | null, sessionId: string | null, meta: RequestMeta) {
    await this.db`INSERT INTO audit_events ${this.db({
      actorUserId: userId, subjectUserId: userId, sessionId,
      eventType, ipAddress: meta.ip, userAgent: meta.userAgent
    })}`;
  }
}

function publicUser(user: Record<string, unknown>) {
  return { id: user.id, email: user.email, username: user.username };
}

function isUniqueViolation(error: unknown): boolean {
  return !!error && typeof error === "object" && "code" in error && error.code === "23505";
}

export class IdentityError extends Error {
  constructor(public readonly code: string, public readonly status: number) { super(code); }
}

export type RequestMeta = { ip: string; userAgent: string | null };
