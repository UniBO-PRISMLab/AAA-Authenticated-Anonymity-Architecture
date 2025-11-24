# Ansible Playbooks

Here the Ansible playbooks used to configure the backend on the host machines used for the hackathon.

To deploy on your own server create a file `hosts.ini` e.g.,

```yaml
[machines]
129.xxx.xx.xxx ansible_user=ubuntu ansible_password=<ubuntu_password>
...
129.xxx.xx.xxx ansible_user=ubuntu ansible_password=<ubuntu_password>
```

Add hosts to known hosts if necessary `ssh-keyscan -H 129.xxx.xx.xxx >> ~/.ssh/known_hosts`.

Hosts are required to have Docker and Git **pre-installed**.

## Deploy

In your active shell export the following environment variables

```bash
export GHCR_USERNAME="<your_gh_username>"
export GH_TOKEN="ghp_token"
export DB_USER_CONTENT="<a_db_user>"
export DB_PASSWORD_CONTENT="<a_db_password>"
export DB_NAME_CONTENT="<a_db_name>"
```

`GHCR_USERNAME` should be your GitHub username. `GH_TOKEN` should be a valid personal access token with grants to clone repository and pull images from UniboPRISMLab GitHub container registry. `DB_*_CONTENT` variables are needed to initialize a PostgreSQL database on the host.

Supposing your hosts are in `~/hosts.ini`, the command

```bash
ansible-playbook -i ~/hosts.ini deploy.yaml
```

will setup the host and start the backend.

## Produce Identities

The command

```bash
ansible-playbook -i ~/hosts.ini identities.yaml
```

will generate 100 public identities (PIDs) and 12 anonymous identities (SIDs). All 100 public identities are copied to `/home/student/public-identities.txt` and 3 anonymous identities are copied to `/home/student/anonymous-identities.txt`.

Hackathon goal is to discover the mapping between a public identity and an anonymous identity.

### Cleanup

To cleanup your hosts run

```bash
ansible-playbook -i ~/hosts.ini cleanup.yaml
```
