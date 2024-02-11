<div>
  <img width="64" height="70" align="left" src="./glowstone.png" alt="Glowstone"/>
  <br>
  <h1>Glowstone</h1>
  <p>Glowstone is a Golang library for reading and writing Minecraft Bedrock Edition addon.</p>
</div>

  # Setup
  To enable multiple workspace of go, you need to create a file `go.work` in your workspace.

  ```
  $ go work init
  ```

  Let's say you want to create go workspace at `./scripts/dummy_entity`
  ```
  $ mkdir ./scripts/dummy_entity
  $ cd ./scripts/dummy_entity
  $ go mod init dummy_entity
  $ go get github.com/respectZ/glowstone
  ```

  Create `main.go` in the `./scripts/dummy_entity` directory with the following contents:
  ```go
  package main

  import (
    "github.com/respectZ/glowstone"
  )

  func main() {
    project := glowstone.NewProject()

    // Set your RP and BP path.
    // Below is the default value.
    project.BP.Path = filepath.Join("packs", "BP")
    project.RP.Path = filepath.Join("packs", "RP")

    // Do something here before saving

    project.Save()
  }
  ```

  Add the project into `go.work`

  ```
  <!-- Back to your workspace -->
  $ cd ../../
  $ go work use ./scripts/dummy_entity
  ```

  To execute your scripts from workspace:

  ```
  $ go run ./scripts/dummy_entity/
  ```

  # Example
  ```go
  package main

  import (
    "path/filepath"
    "github.com/respectZ/glowstone"
    component "github.com/respectZ/glowstone/bp/entity/component"
  )

  func main() {
    // Create a project workspace.
    project := glowstone.NewProject()

    // Set your RP and BP path.
    // Below is the default value.
    project.BP.Path = filepath.Join("packs", "BP")
    project.RP.Path = filepath.Join("packs", "RP")

    // Preload BP and RP packs from the path.
    // Warning: this is unstable and may cause some errors for some compabilities.
    // This isn't necessary if you didn't need any read operation.
    project.Preload()

    NewEntity(project, "glowstone:chair")

    project.Save()
  }

  func NewEntity(project *glowstone.Project, identifier string) {
    bpFile, rpFile := project.NewEntity(identifier)
    bp := bpFile.Data
    rp := rpFile.Data

    rp.Entity.Description.Textures.Add("default", "textures/items/blaze_rod")

    bp.Entity.Description.IsSpawnable = false
	  bp.Entity.ComponentGroups.Add("despawn", &component.InstantDespawn{})

    ev := bp.Entity.Events.New("glowstone:despawn")
	  ev.Add.ComponentGroups.Add("despawn")
  }
  ```

  # Todo
- BP
  - [x] Animation
  - [x] Animation Controller
  - [x] Block
  - [x] Entity
  - [x] Item
  - [x] Loot Table
  - [x] Recipe
  - [ ] Spawn Rule
  - [ ] Structure
- RP
  - [x] Animation Controller
  - [x] Animation
  - [x] Attachable
  - [x] Entity
  - [x] Fog
  - [ ] Material
  - [x] Model
  - [x] Particle
  - [x] Render Controller
  - [x] Item Texture
  - [ ] sounds.json
  - [x] sound_definitions.json
  - [x] blocks.json
  - [x] terrain_textures.json

- Misc
  - [x] Make a field for saving filename, since it will duplicate if thee filename doesn't match with the identifier.
  - [x] Rework project fields, ```glowstone.NewBPAnimationController -> glowstone.BP.AnimationController.New("asdf")```
  - [ ] Remove unnecessary pointer
  - [ ] Add project setting to minify JSON
  - [ ] Add format version setting, especially for item, block, and entity

- Util
  - [ ] BBModel Parser 
- [ ] Documentation

- Format Version
  - [ ] Item
  - [ ] Block
  - [ ] Entity (RP)
  - [ ] Blocks (RP)
  - [ ] Attachable (RP)