package main

// Material type constants
const (
	MaterialTypeDrive   = "drive"
	MaterialTypeLink    = "link"
	MaterialTypeYoutube = "youtube"
	MaterialTypeForm    = "form"
	MaterialTypeUnknown = "unknown"
)

// Material represents a union type for different material types
type Material struct {
	DriveFile    *DriveFileMaterial    `json:"driveFile,omitempty"`
	Link         *LinkMaterial         `json:"link,omitempty"`
	YoutubeVideo *YoutubeVideoMaterial `json:"youtubeVideo,omitempty"`
	Form         *FormMaterial         `json:"form,omitempty"`
}

// DriveFileMaterial - nested driveFile.driveFile structure
type DriveFileMaterial struct {
	DriveFile DriveFileInner `json:"driveFile"`
	ShareMode string         `json:"shareMode"`
}

type DriveFileInner struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	AlternateLink string `json:"alternateLink"`
}

type LinkMaterial struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

type YoutubeVideoMaterial struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	AlternateLink string `json:"alternateLink"`
}

type FormMaterial struct {
	FormURL     string `json:"formUrl"`
	ResponseURL string `json:"responseUrl"`
	Title       string `json:"title"`
}

// Type returns the material type constant
func (m *Material) Type() string {
	switch {
	case m.DriveFile != nil:
		return MaterialTypeDrive
	case m.Link != nil:
		return MaterialTypeLink
	case m.YoutubeVideo != nil:
		return MaterialTypeYoutube
	case m.Form != nil:
		return MaterialTypeForm
	default:
		return MaterialTypeUnknown
	}
}

// GetTitle returns the title of the material
func (m *Material) GetTitle() string {
	switch {
	case m.DriveFile != nil:
		return m.DriveFile.DriveFile.Title
	case m.Link != nil:
		return m.Link.Title
	case m.YoutubeVideo != nil:
		return m.YoutubeVideo.Title
	case m.Form != nil:
		return m.Form.Title
	default:
		return ""
	}
}

// GetURL returns the URL/link of the material
func (m *Material) GetURL() string {
	switch {
	case m.DriveFile != nil:
		return m.DriveFile.DriveFile.AlternateLink
	case m.Link != nil:
		return m.Link.URL
	case m.YoutubeVideo != nil:
		return m.YoutubeVideo.AlternateLink
	case m.Form != nil:
		return m.Form.FormURL
	default:
		return ""
	}
}

// GetID returns the ID of the material (if applicable)
func (m *Material) GetID() string {
	switch {
	case m.DriveFile != nil:
		return m.DriveFile.DriveFile.ID
	case m.YoutubeVideo != nil:
		return m.YoutubeVideo.ID
	default:
		return ""
	}
}

// Implement list.Item for Material
func (m Material) FilterValue() string { return m.GetTitle() }
func (m Material) Title() string       { return m.GetTitle() }
func (m Material) Description() string { return "[" + m.Type() + "] " + m.GetURL() }
