# Robot task for kiwi.com
The application is a simulation of a toy robot moving on a square
tabletop, of dimensions 5 units x 5 units. There are no other
obstructions on the table surface. The robot is free to roam around
the surface of the table, but must be prevented from falling to
destruction. Any movement that would result in the robot falling from
the table must be prevented, however further valid movement commands
must still be allowed.

## Quickstart
Requirements:
- Docker

How to build and run app:
1. Clone the repo
2. Build the image with command `docker build -t kiwi-robot .`
3. Run container with command `docker run kiwi-robot PLACE 0,0,EAST MOVE REPORT`
4. You must see output in STDOUT: `1,0,EAST`
5. Also it is possible to run help with command: `docker run kiwi-robot -h`

## Examples
- `PLACE 0,0,EAST`
    * result: ` `
- `PLACE 0,0,EAST REPORT`
    * result: `0,0,EAST`
- `PLACE 0,0,EAST RIGHT MOVE LEFT MOVE REPORT`
    * result: `1,0,EAST`
- `LEFT MOVE RIGHT REPORT PLACE 2,3,WEST MOVE REPORT`
    * result: `1,3,WEST`
- `PLACE 0,0,EAST MOVE PLACE 5,3,WEST REPORT`
    * result: `1,0,EAST`