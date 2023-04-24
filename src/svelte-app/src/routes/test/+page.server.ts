import type { PageLoadEvent } from "./$types";

export async function load({ parent }: PageLoadEvent) {
  const { uid } = await parent(); // comes from +layout.server.ts
  return { uid };
}