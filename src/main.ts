// main.ts

import {
  applyTransform,
  compareVersions,
  getGitHubLatestVersion,
  getText,
} from "./common.ts";

type AppConfig = {
  name: string;
  installed: string;
  type: "html" | "github";
  url?: string;
  selector?: string;
  transform?: string;
  repo?: string;
};

const apps: AppConfig[] = JSON.parse(
  await Deno.readTextFile("apps.json")
);

for (const app of apps) {
  let latest: string | null = null;

  try {
    if (app.type === "html") {
      const raw = await getText(app.url!, app.selector!);
      latest = applyTransform(raw, app.transform);
    }

    if (app.type === "github") {
      latest = await getGitHubLatestVersion(app.repo!);
    }

    if (!latest) {
      console.log(`X ${app.name}: failed to fetch latest version`);
      continue;
    }

    const result = compareVersions(latest, app.installed);

    if (result > 0) {
      console.log(
        `> ${app.name}: update available (${app.installed} → ${latest})`
      );
    } else {
      console.log(`  ${app.name}: up to date (${app.installed})`);
    }
  } catch (err) {
    console.log(`! ${app.name}: error`, err);
  }
}
