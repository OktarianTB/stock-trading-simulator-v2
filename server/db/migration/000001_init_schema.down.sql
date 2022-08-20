ALTER TABLE IF EXISTS "transactions" DROP CONSTRAINT IF EXISTS "transactions_username_fkey";
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS users;