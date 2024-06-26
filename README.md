# gitpets

### Virtual pets for your README!

made with :heart: for [Beaverhacks Spring 2024](https://spring-2024-retro.devpost.com/?ref_feature=challenge&ref_medium=your-open-hackathons&ref_content=Submissions+open)

## License

[![License](https://img.shields.io/github/license/Ileriayo/markdown-badges?style=for-the-badge)](./LICENSE)

## Usage

<picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://gitpets.fly.dev/api?username=cassiusfive&petname=lorem+ipsum&theme=dark&species=fox">
  <source media="(prefers-color-scheme: light)" srcset="https://gitpets.fly.dev/api?username=cassiusfive&petname=lorem+ipsum&theme=light&species=fox">
  <img alt="Shows a gitpet" src="https://gitpets.fly.dev/api?username=cassiusfive&petname=lorem+ipsum&theme=dark&species=fox" align="left" width="200px">
</picture>

Copy and paste the following snippet into your README

```md
<picture>
  <source media="(prefers-color-scheme: dark)" srcset="https://gitpets.fly.dev/api?username=cassiusfive&petname=lorem+ipsum&theme=dark&species=fox">
  <source media="(prefers-color-scheme: light)" srcset="https://gitpets.fly.dev/api?username=cassiusfive&petname=lorem+ipsum&theme=light&species=fox">
  <img alt="Shows a gitpet" src="https://gitpets.fly.dev/api?username=cassiusfive&petname=lorem+ipsum&theme=dark&species=fox" width="200px">
</picture>

```

<br clear="both"/>

### API Reference
`https://gitpets.fly.dev/api`
| Parameter      | Required | Description                | Options          |
|----------------|----------|----------------------------|------------------|
| **`username`** | Yes      | Github username            | N/A              |
| **`petname`**  | Yes      | Desired name of pet        | N/A              |
| **`species`**  | Yes      | Desired species of pet     | fox, wolf, bunny |
| **`theme`**    | No       | The theme for UI           | light, dark      |


> [!NOTE]
> Pets will gain experience based on your commits, merged PRs and unique repos you've contributed to.

## Development

1. Clone the repo

```sh
git clone https://github.com/cassiusfive/gitpets.git
```

2. Get environment variables

```sh
cp .env.example .env
```

> [!IMPORTANT]
> You'll need to generate a Github [personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-fine-grained-personal-access-token)
> to track activity. It does not need any special permissions.

3. Start the project

```sh
docker compose up -d
```
