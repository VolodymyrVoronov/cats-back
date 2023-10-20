-- CreateTable
CREATE TABLE "Cat" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "createdAt" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" DATETIME NOT NULL,
    "name" TEXT NOT NULL,
    "age" INTEGER NOT NULL,
    "breed" TEXT NOT NULL,
    "photo" TEXT NOT NULL,
    "diseases" TEXT NOT NULL,
    "marked" BOOLEAN NOT NULL DEFAULT false,
    "insurance" BOOLEAN NOT NULL DEFAULT false,
    "information" TEXT NOT NULL,
    "alive" BOOLEAN NOT NULL DEFAULT true,
    "dead" BOOLEAN NOT NULL DEFAULT false
);
