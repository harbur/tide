# Tide

Tide is a standalone CLI tool to render Kubernetes packages with the Helm format.

It is a stateless tool, which uses the current directory to render the templates
and either outputs them for further usage (e.g. pipeline with kubectl or other
scripting methods) or installs them directly to Kubernetes.

# Usage

Checkout your favorite Chart repository, go inside the directory and start some services:

```
git clone https://github.com/kubernetes/charts
# Create an empty values.toml file
touch alpine/values.toml mysql/values.toml
# Start Alpine Chart
tide up alpine
# Stop Alpine Chart
tide down alpine
# Watch Alpine for Changes and apply them (Best used with Deployment objects to use rolling-update feature)
tide up -w alpine
# Watch and Delete Alpine Chart (Best used for development of Charts)
tide up -wd alpine
# Start multiple Charts on one command
tide up alpine mysql
```
