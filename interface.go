package qb

import "image/color"

// VoxelGrid is a 3D grid of volumetric pixels.
// Each pixel had a color, represented as a color.Color.
type VoxelGrid interface {
	// ColorModel returns the VoxelGrid's color model.
	ColorModel() color.Model

	// Returns the smallest cube that fits the entire grid.
	// Trying to read voxels outside this region will result in color.Transparent.
	// There is no origin, the point (0, 0, 0) may not even be within the bounds.
	Bounds() Cube

	// Get returns the color of the voxel at (x, y, z).
	// When reading outside the model bounds, the output color will be transparent.
	Get(x, y, z int) color.Color
}

// To save memory, a VoxelGrid can be split into several smaller chunks.
// The containing structure can still be queried like a regular model.
type VoxelStitch interface {
	VoxelGrid

	// Partitions returns the individual parts of the model.
	// The coordinate space will be the same for the output stitch,
	// meaning (10, 10, 10) on the parent will also be (10, 10, 10) on any child.
	// VoxelStitches can be nested.
	Parts() []VoxelGrid
}

// VoxelCanvas is a VoxelGrid, that supports redefining its voxels.
type VoxelCanvas interface {
	VoxelGrid

	// Plotting a voxel outside the grid bounds will be ignored.
	// It is never guaranteed that a written voxel can be read back,
	// due to VoxelStiches potentially containing unmapped space in their bounds.
	Set(x, y, z int, voxel color.Color)
}

type Point struct {
	X, Y, Z int
}

type Cube struct {
	Min, Max Point
}

func (c Cube) Contains(p Point) bool {
	return p.X >= c.Min.X && p.X < c.Max.X &&
		p.Y >= c.Min.Y && p.Y < c.Max.Y &&
		p.Z >= c.Min.Z && p.Z < c.Max.Z
}
