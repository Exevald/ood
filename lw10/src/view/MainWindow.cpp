#include "MainWindow.h"
#include "CanvasWidget.h"

#include <QColorDialog>
#include <QFileDialog>
#include <QMenuBar>
#include <QToolBar>
#include <QToolButton>

MainWindow::MainWindow(ViewModel* vm, QWidget* parent)
	: QMainWindow(parent)
	, m_vm(vm)
{
	auto* canvas = new CanvasWidget(m_vm, this);
	setCentralWidget(canvas);

	CreateToolbar();
	CreateMenu();

	setWindowTitle("Web editor");
	resize(1024, 768);
}

void MainWindow::CreateToolbar()
{
	QToolBar* toolbar = addToolBar("Tools");

	auto* shapesButton = new QToolButton(this);
	shapesButton->setText("Shapes");
	shapesButton->setToolTip("Insert Shape");
	shapesButton->setPopupMode(QToolButton::InstantPopup);

	auto* shapesMenu = new QMenu(shapesButton);
	const QAction* actRect = shapesMenu->addAction("Rectangle");
	connect(actRect, &QAction::triggered, m_vm, &ViewModel::AddRect);

	const QAction* actTriangle = shapesMenu->addAction("Triangle");
	connect(actTriangle, &QAction::triggered, m_vm, &ViewModel::AddTriangle);

	const QAction* actEllipse = shapesMenu->addAction("Ellipse");
	connect(actEllipse, &QAction::triggered, m_vm, &ViewModel::AddEllipse);

	shapesButton->setMenu(shapesMenu);
	toolbar->addWidget(shapesButton);

	const QAction* actImage = toolbar->addAction("Image");
	connect(actImage, &QAction::triggered, this, &MainWindow::OnAddImage);

	toolbar->addSeparator();

	QAction* actGroup = toolbar->addAction("Group");
	actGroup->setShortcut(QKeySequence(Qt::CTRL + Qt::Key_G));
	connect(actGroup, &QAction::triggered, m_vm, &ViewModel::GroupSelected);

	QAction* actUngroup = toolbar->addAction("Ungroup");
	actUngroup->setShortcut(QKeySequence(Qt::CTRL + Qt::SHIFT + Qt::Key_G));
	connect(actUngroup, &QAction::triggered, m_vm, &ViewModel::UngroupSelected);

	toolbar->addSeparator();

	QAction* actUndo = toolbar->addAction("Undo");
	actUndo->setShortcut(QKeySequence::Undo);
	connect(actUndo, &QAction::triggered, m_vm, &ViewModel::Undo);

	QAction* actRedo = toolbar->addAction("Redo");
	actRedo->setShortcut(QKeySequence::Redo);
	connect(actRedo, &QAction::triggered, m_vm, &ViewModel::Redo);

	toolbar->addSeparator();

	QAction* actCopy = toolbar->addAction("Copy");
	actCopy->setShortcut(QKeySequence::Copy);
	connect(actCopy, &QAction::triggered, m_vm, &ViewModel::Copy);

	QAction* actPaste = toolbar->addAction("Paste");
	actPaste->setShortcut(QKeySequence::Paste);
	connect(actPaste, &QAction::triggered, m_vm, &ViewModel::Paste);

	QAction* actDelete = toolbar->addAction("Delete");
	actDelete->setShortcut(QKeySequence::Delete);
	connect(actDelete, &QAction::triggered, m_vm, &ViewModel::DeleteSelected);

	toolbar->addSeparator();

	const QAction* actColor = toolbar->addAction("Color");
	connect(actColor, &QAction::triggered, this, [this] {
		if (const QColor color = QColorDialog::getColor(Qt::white, this, "Select Color");
			color.isValid())
		{
			m_vm->SetColorForSelected(color.rgba());
		}
	});
}

void MainWindow::CreateMenu()
{
	QMenu* fileMenu = menuBar()->addMenu("File");
	const QAction* actSave = fileMenu->addAction("Save As...");
	connect(actSave, &QAction::triggered, this, &MainWindow::OnSave);

	const QAction* actLoad = fileMenu->addAction("Open...");
	connect(actLoad, &QAction::triggered, this, &MainWindow::OnLoad);
}

void MainWindow::OnAddImage()
{
	if (const QString path = QFileDialog::getOpenFileName(
			this, "Select Image", "", "Images (*.png *.jpg *.jpeg *.bmp)");
		!path.isEmpty())
	{
		m_vm->AddImage(path);
	}
}

void MainWindow::OnSave()
{
	if (const QString path = QFileDialog::getSaveFileName(
			this, "Save Document", "", "JSON (*.json)");
		!path.isEmpty())
	{
		m_vm->Save(path);
	}
}

void MainWindow::OnLoad()
{
	if (const QString path = QFileDialog::getOpenFileName(
			this, "Open Document", "", "JSON (*.json)");
		!path.isEmpty())
	{
		m_vm->Load(path);
	}
}