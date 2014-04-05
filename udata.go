package oden

import "github.com/robertkrimen/otto"

// Data for X, Y cordinates, width and height and angle for rotation
type TransformComponent struct {
	Component
	X, Y  int32
	W, H  int32
	Angle int32
}

func NewTransformComponent(x, y int32) *TransformComponent {
	return &TransformComponent{
		X:     x,
		Y:     y,
		Angle: 0,
	}
}

// The camera used by rendersystem
type CameraComponent struct {
	Component
}

func NewCameraComponent() *CameraComponent {
	return &CameraComponent{}
}

// Data for textures used for rendering
type SpriteComponent struct {
	Component
	Texture Texture2D
}

func NewSpriteComponent() *SpriteComponent {
	return &SpriteComponent{}
}

// Data for the Gel script system
type ScriptComponent struct {
	Component
	// Script loaded either manually or by base.Assets
	Src string
	// So we wont need to run the same script again and again
	object *otto.Object
}

func NewScriptComponent() *ScriptComponent {
	return &ScriptComponent{}
}
