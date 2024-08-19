# gcp-secrets-loader
Automate batch loading of secrets into GCP.

### Usage

#### Authentication

It will automatically pick up `GOOGLE_APPLICATION_DEFAULT_CREDENTIALS`.
Manually passing GCP Service Account key files is not yet supported.

#### Execution

Example:

```bash
go build .

./gcp-secrets-loader \
--path ./secrets/dev_project.csv \     
--project "projects/acmecorp_dev_sandbox" \  
--create                                   
```

Explanation of flags:

| Flag               | Description                                  |
| ------------------ | -------------------------------------------- |
| `--path STRING`    | Path to CSV file, see format below           |
| `--project STRING` | Name of the project in which to load secrets |
| `--create`         | Whether or not to create the secrets         |


#### CSV Format

A secrets CSV file is expected to follow this format:

```
secret_id,secret_value
```

where `secret_id` is the ID of the secret (i.e. not the fully qualified secret URI) and `secret_value` is the plaintext value of the secret.