package entity

import (
	"swordmaster/models"
	"swordmaster/pkg/window"
	"swordmaster/store"

	"github.com/go-gl/glfw/v3.3/glfw"
	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/tfriedel6/canvas"
)

type Player struct {
	position glm.Vec2
	size     float64
	name     string
}

func NewPlayer(name string, x, y float64, s float64) *Player {
	return &Player{
		name: name,
		position: glm.Vec2{
			x, y,
		},
		size: s,
	}
}

func (p *Player) Setup(w *window.Window) {
	w.KB.AddListener(glfw.KeyW, func() {
		p.position = p.position.Add(glm.Vec2{0, -1})
	})
	w.KB.AddListener(glfw.KeyS, func() {
		p.position = p.position.Add(glm.Vec2{0, 1})
	})
	w.KB.AddListener(glfw.KeyA, func() {
		p.position = p.position.Add(glm.Vec2{-1, 0})
	})
	w.KB.AddListener(glfw.KeyD, func() {
		p.position = p.position.Add(glm.Vec2{1, 0})
	})
}

func (p *Player) Draw(cv *canvas.Canvas, w, h float64) {
	cv.SetFillStyle("#00F")
	if store.GetLink() != nil {
		store.GetLink().Broadcast(&models.Message{
			Kind: "POS",
			Name: p.name,
			Data: []float64{p.position.X(), p.position.Y(), 0},
		})
	}
	cv.FillRect(p.position.X(), p.position.Y(), p.size, p.size)
	for _, client := range store.GetClients() {
		cv.FillRect(client.Position[0], client.Position[1], p.size, p.size)
	}
}
