# API Design Doc

## Window

```go
ui := termui.Create(tc Color, bg Color)
```

- Creates main termui window
- this window is a termui container that scales automatically to fit the size of the terminal
- the window must only have one child container that is also a \*Page

## Pages

```go
page := termui.NewPage(tc Color, bg Color, m ContainerMode)

// change to page in ui
ui.SetPage(page)
```

- Pages in termui are just containers that contain all the other containers and content that should be shown on one page
- This allows for easy routing between pages

## Containers

```go
container := termui.NewContainer(tc Color, bg Color, m ContainerMode)
```

- Containers are what every element in your ui is made of
- Can contain other child containers or have their own content depending on mode

### Container Modes

- **BLOCK**

  - cannot contain child containers
  - used for content (text)
  - content can be vertically and horizontally aligned

- **FLEX**

  - can only contain child containers (no content)
  - children can be position in the same way as css flex elements
  - used for simple child positioning

- **GRID**
  - can only contain child containers (no content)
  - number of columns and rows can be set
  - children are positioned along grid lines
  - used for complex layouts
