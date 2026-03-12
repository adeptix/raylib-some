package snowflake

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	optionDepth      = 5
	optionLinesCount = 3
	optionRotateAng  = float64(math.Pi / 2)

	drawCallsCount = 0

	baseColor = rl.Red
	hsvColor  = rl.ColorToHSV(baseColor)
)

//
// 4 + 4*4 + 4*4*4 + 4*4*4*4

func InitWindow() func() {
	rl.InitWindow(800, 600, "snowflake")
	rl.SetTargetFPS(60)

	return func() {
		rl.CloseWindow()
	}
}

func UpdateFunc() {
	const (
		lineLength     = 150
		rotateAngSpeed = math.Pi / 2
	)

	rl.ClearBackground(rl.White)
	rl.BeginDrawing()

	dt := rl.GetFrameTime()

	btnPos := rl.NewVector2(20, 20)
	var btnRec rl.Rectangle

	btnRec = createButton(btnPos, "- depth", func() {
		if optionDepth > 1 {
			optionDepth--
		}
	})
	btnPos.X += btnRec.Width + 20

	btnRec = createButton(btnPos, "+ depth", func() {
		if optionDepth < 10 {
			optionDepth++
		}
	})
	btnPos.X += btnRec.Width + 20*2

	btnRec = createButton(btnPos, "- lines", func() {
		if optionLinesCount > 1 {
			optionLinesCount--
		}
	})
	btnPos.X += btnRec.Width + 20

	btnRec = createButton(btnPos, "+ lines", func() {
		if optionLinesCount < 10 {
			optionLinesCount++
		}
	})
	btnPos.X += btnRec.Width + 20*2

	btnRec = createButton(btnPos, "- ang", func() {
		optionRotateAng -= rotateAngSpeed * float64(dt)
	}, true)
	btnPos.X += btnRec.Width + 20

	btnRec = createButton(btnPos, "+ ang", func() {
		optionRotateAng += rotateAngSpeed * float64(dt)
	}, true)
	btnPos.X += btnRec.Width + 20*2

	drawCallsCount = 0

	//drawSnowflake(rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) / 2}, optionLinesCount, lineLength, hsvColor, optionDepth)
	drawSnowflakeV2(rl.Vector2{X: float32(rl.GetScreenWidth()) / 2, Y: float32(rl.GetScreenHeight()) / 2}, optionLinesCount, lineLength, hsvColor, optionDepth)

	rl.DrawText(fmt.Sprintf("%d", drawCallsCount), int32(rl.GetScreenWidth()-60), 20, 16, rl.Blue)

	rl.EndDrawing()
}

func drawSnowflake(center rl.Vector2, linesCount int, lineLength float32, hsvColor rl.Vector3, depth int) {
	if depth <= 0 {
		return
	}

	clr := rl.ColorFromHSV(hsvColor.X, hsvColor.Y, hsvColor.Z)

	for i := range linesCount {
		ang := 2*math.Pi/float64(linesCount)*float64(i) + optionRotateAng

		dy := math.Sin(ang) * float64(lineLength)
		dx := math.Cos(ang) * float64(lineLength)

		endPos := rl.NewVector2(center.X+float32(dx), center.Y-float32(dy))

		rl.DrawLineEx(center, endPos, 4, clr)
		drawCallsCount++

		drawSnowflake(endPos, linesCount, lineLength/3, rl.NewVector3(hsvColor.X-60, hsvColor.Y, hsvColor.Z), depth-1)
	}
}

// for loop + calc sin/cos only once
func drawSnowflakeV2(center rl.Vector2, linesCount int, lineLength float32, hsvColor rl.Vector3, depth int) {
	cosSins := make([]rl.Vector2, linesCount)

	for i := range linesCount {
		ang := 2*math.Pi/float64(linesCount)*float64(i) + optionRotateAng
		cosSins[i].X = float32(math.Cos(ang))
		cosSins[i].Y = float32(math.Sin(ang))
	}

	verts := []rl.Vector2{center}
	currentLength := lineLength
	var nextVerts []rl.Vector2

	for d := range depth {
		clr := rl.ColorFromHSV(hsvColor.X, hsvColor.Y, hsvColor.Z)
		hasNext := d < depth-1

		if hasNext {
			nextVerts = make([]rl.Vector2, 0, cap(verts)*linesCount)
		}

		for _, vert := range verts {
			for _, cosSin := range cosSins {
				endVert := rl.NewVector2(vert.X+cosSin.X*currentLength, vert.Y+cosSin.Y*currentLength)
				rl.DrawLineEx(vert, endVert, 4, clr)
				drawCallsCount++

				if hasNext {
					nextVerts = append(nextVerts, endVert)
				}
			}
		}

		verts = nextVerts
		currentLength /= 3
		hsvColor.X -= 60
	}
}

// draw with rl.DrawLineStrip ??
func drawSnowflakeV3(center rl.Vector2, linesCount int, lineLength float32, hsvColor rl.Vector3, depth int) {

}

func createButton(pos rl.Vector2, text string, callback func(), onMouseDown ...bool) rl.Rectangle {
	const (
		margin   = 10
		fontSize = 14
		minSize  = 30
	)

	textSize := rl.MeasureText(text, fontSize)
	btnSizeX := textSize + margin
	btnSizeY := int32(fontSize + margin)

	btnSizeX = max(btnSizeX, minSize)
	btnSizeY = max(btnSizeY, minSize)

	rec := rl.NewRectangle(pos.X, pos.Y, float32(btnSizeX), float32(btnSizeY))

	mousePos := rl.GetMousePosition()
	isHovered := rl.CheckCollisionPointRec(mousePos, rec)

	if callback != nil && isHovered {
		isPressed := false
		if len(onMouseDown) == 1 && onMouseDown[0] {
			isPressed = rl.IsMouseButtonDown(rl.MouseButtonLeft)
		} else {
			isPressed = rl.IsMouseButtonPressed(rl.MouseButtonLeft)
		}

		if isPressed {
			callback()
		}
	}

	if isHovered {
		rl.DrawRectangleRec(rec, rl.RayWhite)
	}

	rl.DrawRectangleLinesEx(rec, 2, rl.Gray)
	rl.DrawText(text, int32(rec.X)+btnSizeX/2-textSize/2, int32(rec.Y)+btnSizeY/2-fontSize/2, fontSize, rl.Black)

	return rec
}
