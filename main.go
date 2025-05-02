package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// NewSpacer: 빈 공간을 생성하는 함수
func NewSpacer() *fyne.Container {
	return container.NewWithoutLayout(layout.NewSpacer())
}

// createColor: RGBA 값을 간단히 입력하여 color.Color를 생성하는 헬퍼 함수
func createColor(r, g, b, a uint8) color.Color {
	return color.RGBA{R: r, G: g, B: b, A: a}
}

// createVerticalLayout: 수직 레이아웃 생성
func createVerticalLayout() *fyne.Container {
	textBox1 := widget.NewEntry()
	textBox1.SetPlaceHolder("텍스트 박스 1")

	button1 := widget.NewButton("버튼 1", func() {})

	textBox2 := widget.NewEntry()
	textBox2.SetPlaceHolder("텍스트 박스 2")

	button2 := widget.NewButton("버튼 2", func() {})

	return container.NewVBox(
		widget.NewLabel("수직 레이아웃"),
		textBox1,
		button1,
		textBox2,
		button2,
	)
}

// createHorizontalLayout: 수평 레이아웃 생성
func createHorizontalLayout() *fyne.Container {
	textBox2 := widget.NewEntry()
	textBox2.SetPlaceHolder("텍스트 박스 2")

	button2 := widget.NewButton("버튼 2", func() {})

	return container.NewMax(
		container.NewAppTabs(
			container.NewTabItem("텍스트 박스", container.NewVBox(
				widget.NewLabel("텍스트 입력"),
				textBox2,
				widget.NewButton("확인", func() {
					fmt.Println("텍스트 입력 확인: ", textBox2.Text)
				}),
			)),
			container.NewTabItem("버튼", container.NewVBox(
				widget.NewLabel("버튼 동작"),
				button2,
				widget.NewButton("추가 버튼", func() {
					fmt.Println("추가 버튼 클릭")
				}),
			)),
		),
	)
}

// createGridLayout: 그리드 레이아웃 생성
func createGridLayout() *fyne.Container {
	textBox1 := widget.NewEntry()
	textBox1.SetPlaceHolder("텍스트 박스 1")

	button1 := widget.NewButton("버튼 1", func() {})

	textBox2 := widget.NewEntry()
	textBox2.SetPlaceHolder("텍스트 박스 2")

	button2 := widget.NewButton("버튼 2", func() {})

	return container.NewGridWithColumns(2,
		widget.NewLabel("그리드 레이아웃"),
		NewSpacer(),
		textBox1,
		button1,
		textBox2,
		button2,
	)
}

// createScrollableList: 스크롤 가능한 목록 생성
func createScrollableList() *fyne.Container {
	listData := []string{"아이템 1", "아이템 2", "아이템 3"}
	list := widget.NewList(
		func() int { return len(listData) },
		func() fyne.CanvasObject { return widget.NewLabel("템플릿") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(listData[i])
		},
	)

	addButton := widget.NewButton("항목 추가", func() {
		listData = append(listData, fmt.Sprintf("아이템 %d", len(listData)+1))
		list.Refresh()
	})

	return container.NewBorder(
		addButton,
		nil,
		nil,
		nil,
		container.NewVScroll(list),
	)
}

// customTheme: 사용자 정의 테마 구조체
type customTheme struct{}

func (c customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		return createColor(40, 44, 52, 255) // 차콜 그레이 배경
	}
	if name == theme.ColorNamePrimary {
		return createColor(0, 123, 255, 255) // 밝은 블루 강조 색상
	}
	if name == theme.ColorNameButton {
		return createColor(60, 179, 113, 255) // 밝은 초록 버튼
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (c customTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (c customTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (c customTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(&customTheme{}) // 사용자 정의 테마 적용
	myWindow := myApp.NewWindow("다양한 레이아웃 예제")

	// 탭 컨테이너로 레이아웃 전환
	tabs := container.NewAppTabs(
		container.NewTabItem("수직", createVerticalLayout()),
		container.NewTabItem("수평", createHorizontalLayout()),
		container.NewTabItem("그리드", createGridLayout()),
		container.NewTabItem("목록", createScrollableList()),
	)

	// 창에 콘텐츠 설정 및 크기 조정
	myWindow.SetContent(tabs)               // tabs 컨테이너를 창의 콘텐츠로 설정합니다.
	myWindow.Resize(fyne.NewSize(400, 300)) // 창의 초기 크기를 설정합니다. 너비 400, 높이 300으로 설정됩니다.
	myWindow.ShowAndRun()                   // 애플리케이션 실행 및 창 표시
}
