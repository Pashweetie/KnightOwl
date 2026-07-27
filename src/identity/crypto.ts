import { createHmac, randomBytes } from "node:crypto";
import { hash, verify, Algorithm } from "@node-rs/argon2";

const argon = {
  algorithm: Algorithm.Argon2id,
  memoryCost: 65536,
  timeCost: 3,
  parallelism: 1,
  outputLen: 32
} as const;

export function newOpaqueToken(): string {
  return randomBytes(32).toString("base64url");
}

export function tokenHash(token: string, key: string): Buffer {
  return createHmac("sha256", key).update(token).digest();
}

export function hashPassword(password: string, pepper: string): Promise<string> {
  return hash(`${password}\u0000${pepper}`, argon);
}

export function verifyPassword(encoded: string, password: string, pepper: string): Promise<boolean> {
  return verify(encoded, `${password}\u0000${pepper}`, argon);
}
