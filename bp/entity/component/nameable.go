package component

// Code generated by scrapper/entityBp.ts;

// [Not a component] Describes the special names for this entity and the events to call when the entity acquires those names
type NameableNameActions struct {
  // List of special names that will cause the events defined in 'on_named' to fire
  NameFilter string `json:"name_filter,omitempty"`
  // Event to be called when this entity acquires the name specified in 'name_filter'
  OnNamed string `json:"on_named,omitempty"`
}

// Allows this entity to be named (e.g. using a name tag).  
type Nameable struct {
  // If true, this entity can be renamed with name tags
  AllowNameTagRenaming bool `json:"allow_name_tag_renaming,omitempty"`
  // If true, the name will always be shown
  AlwaysShow bool `json:"always_show,omitempty"`
  // Trigger to run when the entity gets named
  DefaultTrigger string `json:"default_trigger,omitempty"`
  // Describes the special names for this entity and the events to call when the entity acquires those names
  NameActions NameableNameActions `json:"name_actions,omitempty"`
}