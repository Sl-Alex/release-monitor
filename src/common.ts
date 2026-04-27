// How to turn remote deps into local
// 1. set use_remote to true and use loadModule with your import like it is done for "deno_dom" below. After you run your script, Deno will download dependencies to a vendor folder
// 2. Copy the dependencies to the libs folder
// 3. Turn use_remote back to false to use dependencies you just copied

const useRemote = false;

async function loadModule(modulePath: string, useRemote = false) {
  const resolvedPath = useRemote
    ? `https://deno.land/x/${modulePath}`
    : (() => {
        const base = Deno.env.get("DENO_LIB_PATH")!;

        return new URL(
          `./${modulePath}`,
          `file://${base.replaceAll("\\", "/")}/`
        ).href;
      })();

  return await import(resolvedPath);
}

const { DOMParser } = await loadModule(
  "deno_dom@v0.1.56/deno-dom-wasm.ts",
  useRemote
);

export async function getText(url: string, selector: string) {
  const html = await fetch(url).then((r) => r.text());
  const doc = new DOMParser().parseFromString(html, "text/html");

  return doc?.querySelector(selector)?.textContent?.trim() ?? null;
}

export async function getGitHubLatestVersion(repo: string) {
  const cmd = new Deno.Command("gh", {
    args: ["api", `repos/${repo}/releases/latest`],
    stdout: "piped",
  });

  const output = await cmd.output();
  const text = new TextDecoder().decode(output.stdout);

  const data = JSON.parse(text);
  return data.tag_name ?? null;
}

export function applyTransform(value: string | null, transform?: string) {
  if (!value) return null;
  if (!transform) return value;

  if (transform.startsWith("split:")) {
    const token = transform.slice(6);
    return value.split(token)[1]?.trim() ?? null;
  }

  if (transform.startsWith("regex:")) {
    const pattern = transform.slice(6);
    const match = value.match(new RegExp(pattern));

    return match?.[0] ?? null;
  }

  return value;
}

export function compareVersions(a: string, b: string) {
  const cleanA = a.replace(/^v/, "");
  const cleanB = b.replace(/^v/, "");

  const pa = cleanA.split(".").map(Number);
  const pb = cleanB.split(".").map(Number);

  const max = Math.max(pa.length, pb.length);

  for (let i = 0; i < max; i++) {
    const x = pa[i] ?? 0;
    const y = pb[i] ?? 0;

    if (x > y) return 1;
    if (x < y) return -1;
  }

  return 0;
}
