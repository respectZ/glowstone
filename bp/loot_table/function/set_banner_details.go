package function

// NewSetBannerDetails is the function used to set the banner details of the item.
//
// Example:
//
//  // Default banner, without patterns
//  NewSetBannerDetails("default", "black", nil)
//  // Banner with patterns
//  NewSetBannerDetails("default", "black", &[]map[string]string{
//  	{
//  		"pattern": "circle",
//  		"color":   "white",
//  	},
//  })
//
//  - Available types:
//    - default
//    - illager_captain
//
//  - Available colors:
//    - black
//    - blue
//    - brown
//    - cyan
//    - gray
//    - green
//    - light_blue
//    - light_green
//    - magenta
//    - orange
//    - pink
//    - purple
//    - red
//    - silver
//    - white
//    - yellow
//
//  - Available patterns:
//   - base
//   - border
//   - bricks
//   - circle
//   - creeper
//   - cross
//   - curly_border
//   - diagonal_left
//   - diagonal_right
//   - diagonal_up_left
//   - diagonal_up_right
//   - flower
//   - gradient
//   - gradient_up
//   - half_horizontal
//   - half_horizontal_bottom
//   - half_vertical
//   - half_vertical_right
//   - mojang
//   - piglin
//   - rhombus
//   - skull
//   - small_stripes
//   - square_bottom_left
//   - square_bottom_right
//   - square_top_left
//   - square_top_right
//   - stripe_bottom
//   - stripe_center
//   - stripe_downleft
//   - stripe_downright
//   - stripe_left
//   - stripe_middle
//   - stripe_right
//   - stripe_top
//   - triangle_bottom
//   - triangle_top
//   - triangles_bottom
//   - triangles_top
func NewSetBannerDetails(Type string, baseColor string, patterns *[]map[string]string, conditions ...map[string]interface{}) *map[string]interface{} {
	r := make(map[string]interface{})
	r["function"] = "set_banner_details"
	r["type"] = Type
	r["base_color"] = baseColor
	if patterns != nil {
		r["patterns"] = *patterns
	}
	if len(conditions) > 0 {
		r["conditions"] = make([]map[string]interface{}, 0)
	}
	for _, c := range conditions {
		r["conditions"] = append(r["conditions"].([]map[string]interface{}), c)
	}
	return &r
}

// Enums
const (
	// Type
	SetBannerDetailsTypeDefault        string = "default"
	SetBannerDetailsTypeIllagerCaptain string = "illager_captain"
	// BaseColor
	SetBannerDetailsBaseColorBlack      string = "black"
	SetBannerDetailsBaseColorBlue       string = "blue"
	SetBannerDetailsBaseColorBrown      string = "brown"
	SetBannerDetailsBaseColorCyan       string = "cyan"
	SetBannerDetailsBaseColorGray       string = "gray"
	SetBannerDetailsBaseColorGreen      string = "green"
	SetBannerDetailsBaseColorLightBlue  string = "light_blue"
	SetBannerDetailsBaseColorLightGreen string = "light_green"
	SetBannerDetailsBaseColorMagenta    string = "magenta"
	SetBannerDetailsBaseColorOrange     string = "orange"
	SetBannerDetailsBaseColorPink       string = "pink"
	SetBannerDetailsBaseColorPurple     string = "purple"
	SetBannerDetailsBaseColorRed        string = "red"
	SetBannerDetailsBaseColorSilver     string = "silver"
	SetBannerDetailsBaseColorWhite      string = "white"
	SetBannerDetailsBaseColorYellow     string = "yellow"
	// Pattern
	SetBannerDetailsPatternBase                 string = "base"
	SetBannerDetailsPatternBorder               string = "border"
	SetBannerDetailsPatternBricks               string = "bricks"
	SetBannerDetailsPatternCircle               string = "circle"
	SetBannerDetailsPatternCreeper              string = "creeper"
	SetBannerDetailsPatternCross                string = "cross"
	SetBannerDetailsPatternCurlyBorder          string = "curly_border"
	SetBannerDetailsPatternDiagonalLeft         string = "diagonal_left"
	SetBannerDetailsPatternDiagonalRight        string = "diagonal_right"
	SetBannerDetailsPatternDiagonalUpLeft       string = "diagonal_up_left"
	SetBannerDetailsPatternDiagonalUpRight      string = "diagonal_up_right"
	SetBannerDetailsPatternFlower               string = "flower"
	SetBannerDetailsPatternGradient             string = "gradient"
	SetBannerDetailsPatternGradientUp           string = "gradient_up"
	SetBannerDetailsPatternHalfHorizontal       string = "half_horizontal"
	SetBannerDetailsPatternHalfHorizontalBottom string = "half_horizontal_bottom"
	SetBannerDetailsPatternHalfVertical         string = "half_vertical"
	SetBannerDetailsPatternHalfVerticalRight    string = "half_vertical_right"
	SetBannerDetailsPatternMojang               string = "mojang"
	SetBannerDetailsPatternPiglin               string = "piglin"
	SetBannerDetailsPatternRhombus              string = "rhombus"
	SetBannerDetailsPatternSkull                string = "skull"
	SetBannerDetailsPatternSmallStripes         string = "small_stripes"
	SetBannerDetailsPatternSquareBottomLeft     string = "square_bottom_left"
	SetBannerDetailsPatternSquareBottomRight    string = "square_bottom_right"
	SetBannerDetailsPatternSquareTopLeft        string = "square_top_left"
	SetBannerDetailsPatternSquareTopRight       string = "square_top_right"
	SetBannerDetailsPatternStripeBottom         string = "stripe_bottom"
	SetBannerDetailsPatternStripeCenter         string = "stripe_center"
	SetBannerDetailsPatternStripeDownleft       string = "stripe_downleft"
	SetBannerDetailsPatternStripeDownright      string = "stripe_downright"
	SetBannerDetailsPatternStripeLeft           string = "stripe_left"
	SetBannerDetailsPatternStripeMiddle         string = "stripe_middle"
	SetBannerDetailsPatternStripeRight          string = "stripe_right"
	SetBannerDetailsPatternStripeTop            string = "stripe_top"
	SetBannerDetailsPatternTriangleBottom       string = "triangle_bottom"
	SetBannerDetailsPatternTriangleTop          string = "triangle_top"
	SetBannerDetailsPatternTrianglesBottom      string = "triangles_bottom"
	SetBannerDetailsPatternTrianglesTop         string = "triangles_top"
)
