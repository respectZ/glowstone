<div>
  <img width="64" height="70" align="left" src="./glowstone.png" alt="Glowstone"/>
  <br>
  <h1>Glowstone</h1>
  <p>Glowstone is a Golang library for writing Minecraft Bedrock Edition addon.</p>
</div>

# Todo
- Rework entity, so will be NewEvent, NewProps ...
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
  - [ ] Animation Controller
  - [ ] Animation (most likely no, just the header information)
  - [x] Attachable
  - [x] Entity
  - [ ] Material
  - [ ] Model
  - [ ] Particle
  - [ ] Render Controller
  - [x] Item Texture
  - [ ] sounds.json
  - [x] sound_definitions.json
  - [x] blocks.json
  - [x] terrain_textures.json

- Misc
  - [ ] Rework project fields, ```glowstone.NewBPAnimationController -> glowstone.BP.AnimationController.New("asdf")```
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