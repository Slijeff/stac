## STAC - Simple Tool for Automation Control
Stac is a CI/CD tool written in Go that aims for simplicity.

## Why?
Have you ever started a small demo project and you want to setup a CD pipeline on a new server? Scrolling through Jenkin's long list of installation instructions is just too much for someone looking for a quick and simple solution.

## Setup In 4 Steps
1. Download the binary for your server architecture and prepare a `config.json` file in the same directory as the binary. Fill in the address, ports, stac password, and the log file names.
2. Run the binary with the config file as its argument. e.g.`./stac config.json`. Then go to the serving address and access `/f` for frontend control. Register yor repo using Register Repo (remember to provide the stac password you set in the config file).
3. Setup your server address in the GitHub repository webhook setting.
4. Create a `Stacfile.yaml` in your repo. Commit & Push and you are done!

## Some details
- For `Stacfile.yaml` format, see the `Stacfile_sample.yaml`.
- Each commands block execute in parallel. One potential use case is that one command block can build frontend and one can build the backend.
- View the logs to see command execution results. stacLog is for log associated with stac itself, whereas execLog is the log for command execution.

![441d6d207182639f9baf4965fd8bc19d75b008e0](https://user-images.githubusercontent.com/31459252/234895494-b78e282e-f9cd-4a96-b489-a7ff840b9326.png)
