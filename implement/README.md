# master thesis implementation

## Deployment
### Server
1. `make db`
2. `make`

### Pi
1. `make pi`

## Folder Structure
```
.
├── anc                 # server component for doing ANC.
├── client              # mock meetingbox core code for development.
├── estimat-shift       # server component for estimating shift.
├── estmanc             # perf. test for estimating shift and anc.
├── jammer              # component for jamming in meetingbox, compiles and updloads to arduino.
├── pi                  # meetingbox core code, deploy to raspberry pi.
└── server              # unsealing service providor core code, deploy to server.
```