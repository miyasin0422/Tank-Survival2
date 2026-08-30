components {
  id: "damage"
  component: "/main/enemy/damage.script"
}
embedded_components {
  id: "label"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "color {\n"
  "  y: 0.0\n"
  "  z: 0.0\n"
  "}\n"
  "outline {\n"
  "  x: 1.0\n"
  "}\n"
  "text: \"-10\"\n"
  "font: \"/main/assets/mgen.font\"\n"
  "material: \"/builtins/fonts/label.material\"\n"
  ""
  scale {
    x: 0.5
    y: 0.5
  }
}
