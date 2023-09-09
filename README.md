## Overview
This repository contains all the example microservices we used for our gateway
TestPosterService and TestViewerService are integrated into our testing suite here.
Do not run go build due to dependency conflict between Windows and MacOS.
If you get permission denied error when running binaries, can run chmod +x ./{binary name} (On MacOS)

### Set up
1. Do `cd [directory name]` for the microservice you want to run
2. Run the binary (./{name} for MacOS, ./{name}.exe for Windows)
