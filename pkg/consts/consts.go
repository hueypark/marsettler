package consts

const (
	// TPS is ticks per second.
	TPS int = 60

	// Delta is delta(millie seconds) per tick.
	Delta int = 1000 / TPS

	// Node size.
	NodeSize = 165

	// Node size squared.
	NodeSizeSq = NodeSize * NodeSize

	// Half of node size.
	NodeSizeHalf = NodeSize / 2

	// Half of node size squared.
	NodeSizeHalfSq = NodeSizeHalf * NodeSizeHalf
)
