const url = "http://localhost:8080/data"

var globalData = {}
var globalDict = {}

var i = 0
var ants = []
var step = 0
var antsDic = {}
var start = false
var drawBool = false
var stop = false


function setup() {
    createCanvas(window.innerWidth, window.innerHeight)
    background(0)
    frameRate(30)
    fetch(url)
      .then(res => res.json())
      .then(data => ready(data))
}

function draw() {
    if (start) {
        console.log(step)
        if (step == globalData.Steps.length) {
            stop = true
            window.stop()
        }
        console.log(drawBool)
        console.log(i)

        if (drawBool) {
            drawAnts()
            if (i == 60) {
                drawBool = false
                changeArr()
                console.log("Step: " + step)
                step++
                if (step == globalData.Steps.length) {
                    stop = true
                    window.stop()
                }
                i = 0
                calculatePoints()
            }
        } else {
            if (i == 30) {
                drawBool = true
                i = 0
            }
        }
        i++
    }
}

function drawAnts() {
    background(0)
    drawRoomsRepeat()
    strokeWeight(0)
    fill(255, 204, 0)
    // Reprint each ant at given step
    console.log(step)
    for (const [ant, room]  of Object.entries(globalData.Steps[step])) {
        console.log(antsDic[ant].current[0])
        console.log(antsDic[ant].next[0])

        var x = antsDic[ant].current[0] + antsDic[ant].step[0] * i
        var y = antsDic[ant].current[1] + antsDic[ant].step[1] * i
        circle(x, y, 20)
    }
}

function ready(data) {
    globalData = data
    drawRooms()
}

function drawRoomsRepeat() {
    var maxX = 0
    var maxY = 0
    for (const [name, room] of Object.entries(globalData.Rooms)) {
        if (room.X > maxX) {
            maxX = room.X
        }
        if (room.Y > maxY) {
            maxY = room.Y
        }
    }

    for (const [name, room] of Object.entries(globalData.Rooms)) {
        if (room.X > maxX) {
            maxX = room.X
        }
        if (room.Y > maxY) {
            maxY = room.Y
        }
    }

    maxX += 50
    maxY += 50
    var height = window.innerHeight
    var width = window.innerWidth

    var stepX = width / maxX
    var stepY = height / maxY

    for (const [name, room] of Object.entries(globalData.Links)) {
        stroke('rgb(0,255,0)')
        strokeWeight(5)
        line(globalDict[room.Room1][0], globalDict[room.Room1][1], globalDict[room.Room2][0], globalDict[room.Room2][1])
    }

    strokeWeight(0)

    for (const [name, room] of Object.entries(globalData.Rooms)) {
        if (name == globalData.Start.Name || name==globalData.End.Name) {
            continue
        }
        fill(color('red'))
        // textSize(30)
        // text(room.Name, (room.X+1)*stepX-8, (room.Y+1)*stepY+50)
        fill('rgb(0,255,0)')
        circle((room.X+1)*stepX, (room.Y+1)*stepY, 50)
    }

    // textSize(30)
    fill(color('red'))
    // text(globalData.Start.Name, (globalData.Start.X+1)*stepX-8, (globalData.Start.Y+1)*stepY+50);
    strokeWeight(5)
    circle((globalData.Start.X+1)*stepX, (globalData.Start.Y+1)*stepY, 50)

    strokeWeight(0)
    // textSize(30)
    // text(globalData.End.Name, (globalData.End.X+1)*stepX-8, (globalData.End.Y+1)*stepY+55)
    strokeWeight(5)
    circle((globalData.End.X+1)*stepX, (globalData.End.Y+1)*stepY, 50)

}

function drawRooms() {
    for(var i = 0; i < globalData.Ants; i++) {
        ants.push(globalData.Start.Name);
    }

    var maxX = 0
    var maxY = 0
    for (const [name, room] of Object.entries(globalData.Rooms)) {
        if (room.X > maxX) {
            maxX = room.X
        }
        if (room.Y > maxY) {
            maxY = room.Y
        }
    }

    for (const [name, room] of Object.entries(globalData.Rooms)) {
        if (room.X > maxX) {
            maxX = room.X
        }
        if (room.Y > maxY) {
            maxY = room.Y
        }
    }

    // maxX += 50
    // maxY += 50
    var height = window.innerHeight
    var width = window.innerWidth

    var stepX = width / maxX
    var stepY = height / maxY

    coords = {}
    for (const [name, room] of Object.entries(globalData.Rooms)) {
        coords[name] = [(room.X+1)*stepX, (room.Y+1)*stepY]
    }
    globalDict = coords
    for (const [name, room] of Object.entries(globalData.Links)) {
        stroke('rgb(0,255,0)')
        strokeWeight(5)
        line(coords[room.Room1][0], coords[room.Room1][1], coords[room.Room2][0], coords[room.Room2][1])
    }

    strokeWeight(0)

    for (const [name, room] of Object.entries(globalData.Rooms)) {
        if (name == globalData.Start.Name || name==globalData.End.Name) {
            continue
        }
        fill(color('red'))
        // textSize(30)
        // text(room.Name, (room.X+1)*stepX-8, (room.Y+1)*stepY+50)
        fill('rgb(0,255,0)')
        circle((room.X+1)*stepX, (room.Y+1)*stepY, 50)
    }

    // textSize(30)
    fill(color('red'))
    // text(globalData.Start.Name, (globalData.Start.X+1)*stepX-8, (globalData.Start.Y+1)*stepY+50);
    strokeWeight(5)
    circle((globalData.Start.X+1)*stepX, (globalData.Start.Y+1)*stepY, 50)

    strokeWeight(0)
    // textSize(30)
    // text(globalData.End.Name, (globalData.End.X+1)*stepX-8, (globalData.End.Y+1)*stepY+55)
    strokeWeight(5)
    circle((globalData.End.X+1)*stepX, (globalData.End.Y+1)*stepY, 50)

}

function mouseClicked() {
    calculatePoints()
    start = true
}

function calculatePoints() {
    console.log(step)

    console.log(globalData.Steps[step])

    for ([ant, room] of Object.entries(globalData.Steps[step])) {

        var room1 = ants[ant-1]
        var room2 = room


        var x1 = globalDict[room1][0]
        var y1 = globalDict[room1][1]
        var x2 = globalDict[room2][0]
        var y2 = globalDict[room2][1]

        temp = {}
        temp.current = [x1, y1]
        temp.next = [x2, y2]

        // antsDic[ant].current = [x1, y1]
        // antsDic[ant].next = [x2, y2]

        x_diff = x2-x1
        y_diff = y2-y1

        x_step = x_diff / 60
        y_step = y_diff / 60

        temp.step = [x_step, y_step]

        antsDic[ant] = temp
    
    }
}

function changeArr() {
    for ([ant, room] of Object.entries(globalData.Steps[step])) {
        ants[ant-1] = room
    }
}
  