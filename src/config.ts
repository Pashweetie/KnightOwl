import { z } from "zod";

const schema = z.object({
  NODE_ENV: z.enum(["development", "test", "production"]).default("development"),
  PORT: z.coerce.number().int().min(1).max(65535).default(8787),
  DATABASE_URL: z.string().url(),
  PASSWORD_PEPPER: z.string().min(32),
  SESSION_HASH_KEY: z.string().min(32)
});

export type Config = z.infer<typeof schema>;
export const loadConfig = (input: NodeJS.ProcessEnv = process.env): Config => schema.parse(input);
