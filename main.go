package main

import (
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	board         *Board
	currentPlayer int
	state         int
	cpuMode       bool
}

func (g *Game) init() {
	g.board = NewBoard()
	g.currentPlayer = PlayerBlack
	g.state = GameStateTitle
	g.cpuMode = true
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		mouseX, mouseY := ebiten.CursorPosition()
		row := mouseY / 100
		col := mouseX / 100
		// クリック位置に配置可能なら置いて反転する
		if g.board.canPlace(g.currentPlayer, row, col) {
			g.board.placeStone(g.currentPlayer, row, col)
			g.board.flipStones(g.currentPlayer, row, col)
			g.currentPlayer ^= 1
			// g.board.showHint(g.currentPlayer)

			// CPU モードなら自動配置
			vms := g.board.findValidMoves(g.currentPlayer)
			if len(vms) > 0 && g.cpuMode {
				go func() {
					time.Sleep(1 * time.Second)
					g.board.placeAuto(g.currentPlayer)
					g.currentPlayer ^= 1
				}()
			}
		}
	}
	return nil
}

func (g *Game) DrawBoard(screen *ebiten.Image) {
	// Draw Board
	vector.DrawFilledRect(screen, 0, 0, boardWidth, boardHeight, boardColor, true)
	vector.DrawFilledCircle(screen, 200, 200, 5, boardLineColor, true)
	vector.DrawFilledCircle(screen, 600, 200, 5, boardLineColor, true)
	vector.DrawFilledCircle(screen, 200, 600, 5, boardLineColor, true)
	vector.DrawFilledCircle(screen, 600, 600, 5, boardLineColor, true)

	// Draw Stones
	for row := 0; row < numCellsY; row++ {
		for col := 0; col < numCellsX; col++ {
			vector.StrokeRect(screen, float32(col*100), float32(row*100), 100, 100, 1, boardLineColor, true)
			if (*g.board)[row][col] == 0 {
				vector.DrawFilledCircle(screen, float32((col*100)+50), float32((row*100)+50), stoneRadius, stoneColorBlack, true)
			}
			if (*g.board)[row][col] == 1 {
				vector.DrawFilledCircle(screen, float32((col*100)+50), float32((row*100)+50), stoneRadius, stoneColorWhite, true)
			}
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case GameStateTitle:
		vector.DrawFilledRect(screen, 0, 0, boardWidth, boardHeight, boardColor, true)
		go func() {
			time.Sleep(1 * time.Second)
			g.state = GameStateEnd
		}()
	case GameStatePlay:
		g.DrawBoard(screen)
	case GameStateEnd:
		g.DrawBoard(screen)
		vector.DrawFilledRect(screen, 0, 0, boardWidth, boardHeight, color.RGBA{0x00, 0x00, 0x00, 0xb4}, true)
		ebitenutil.DebugPrint(screen, "Game Over!")
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("リバーシ！")

	g := &Game{}
	g.init()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
