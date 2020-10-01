package recipe

// store implements interface for interacting with recipes database
type store interface {
	create(
		title string,
		preparationTime string,
		serves string,
		ingredients string,
		cost int,
	) (*recipe, error)
}
