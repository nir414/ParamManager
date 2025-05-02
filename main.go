package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// NewSpacer: 빈 공간을 생성하는 함수
func NewSpacer() *fyne.Container {
	return container.NewWithoutLayout(layout.NewSpacer())
}

func main() {
	// 애플리케이션 생성
	myApp := app.New()
	myWindow := myApp.NewWindow("다양한 레이아웃 예제")

	// 텍스트 박스와 버튼 생성
	textBox1 := widget.NewEntry() // 첫 번째 텍스트 박스
	textBox1.SetPlaceHolder("텍스트 박스 1")

	textBox2 := widget.NewEntry() // 두 번째 텍스트 박스
	textBox2.SetPlaceHolder("텍스트 박스 2")

	button1 := widget.NewButton("버튼 1", func() {}) // 첫 번째 버튼
	button2 := widget.NewButton("버튼 2", func() {}) // 두 번째 버튼

	// 수직 레이아웃: 텍스트 박스와 버튼을 수직으로 배치
	verticalLayout := container.NewVBox(
		widget.NewLabel("수직 레이아웃"),
		textBox1,
		button1,
		textBox2,
		button2,
	)

	// 수평 레이아웃을 탭으로 변경하여 더 나은 디자인 제공
	horizontalLayout := container.NewAppTabs(
		container.NewTabItem("텍스트 박스", container.NewVBox(
			widget.NewLabel("텍스트 입력"),
			textBox2, // 텍스트 박스를 직접 사용
			widget.NewButton("확인", func() { // 추가 버튼
				fmt.Println("텍스트 입력 확인: ", textBox2.Text)
			}),
		)),
		container.NewTabItem("버튼", container.NewVBox(
			widget.NewLabel("버튼 동작"),
			button2, // 버튼을 중앙에 배치
			widget.NewButton("추가 버튼", func() { // 또 다른 버튼
				fmt.Println("추가 버튼 클릭")
			}),
		)),
	)

	// 그리드 레이아웃: 텍스트 박스와 버튼을 그리드 형태로 배치
	gridLayout := container.NewGridWithColumns(2,
		widget.NewLabel("그리드 레이아웃"),
		NewSpacer(), // 빈 공간을 생성
		textBox1,
		button1,
		textBox2,
		button2,
	)

	// 스크롤 가능한 목록 생성
	listData := []string{"아이템 1", "아이템 2", "아이템 3"} // 초기 데이터
	list := widget.NewList(
		func() int { return len(listData) },                        // 항목 수 반환
		func() fyne.CanvasObject { return widget.NewLabel("템플릿") }, // 항목 템플릿
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(listData[i]) // 항목 데이터 설정
		},
	)

	// 항목 추가 버튼
	addButton := widget.NewButton("항목 추가", func() {
		listData = append(listData, fmt.Sprintf("아이템 %d", len(listData)+1))
		list.Refresh() // 목록 새로고침
	})

	// 스크롤 가능한 컨테이너 생성
	scrollableLayout := container.NewBorder(
		addButton,                  // 상단에 추가 버튼 배치
		nil,                        // 하단 없음
		nil,                        // 왼쪽 없음
		nil,                        // 오른쪽 없음
		container.NewVScroll(list), // 스크롤 가능한 목록
	)

	// 탭 컨테이너로 레이아웃 전환
	tabs := container.NewAppTabs(
		container.NewTabItem("수직", verticalLayout),
		container.NewTabItem("수평", horizontalLayout),
		container.NewTabItem("그리드", gridLayout),
		container.NewTabItem("목록", scrollableLayout),
	)

	// 창에 콘텐츠 설정 및 크기 조정
	myWindow.SetContent(tabs)               // tabs 컨테이너를 창의 콘텐츠로 설정합니다.
	myWindow.Resize(fyne.NewSize(400, 300)) // 창의 초기 크기를 설정합니다. 너비 400, 높이 300으로 설정됩니다.
	myWindow.ShowAndRun()                   // 애플리케이션 실행 및 창 표시
}
