components {
  id: "upper"
  component: "/main/enemy/enemy_upper.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/main/assets/images/tank/tank.atlas\"\n"
  "default_animation: \"upper_red\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 6.0
    z: 0.9
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "bullet"
  type: "factory"
  data: "prototype: \"/main/power_enemy/power_enemy_bullet.go\"\n"
  "load_dynamically: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "fire"
  type: "sound"
  data: "sound: \"/main/sound/bomb2.ogg\"\n"
  "looping: 0\n"
  "group: \"master\"\n"
  "gain: 1.0\n"
  "pan: 0.0\n"
  "speed: 1.0\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
