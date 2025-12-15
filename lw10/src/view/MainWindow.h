#pragma once

#include "../viewModel/ViewModel.h"

#include <QMainWindow>

class MainWindow final : public QMainWindow
{
	Q_OBJECT

public:
	explicit MainWindow(ViewModel* vm, QWidget* parent = nullptr);

private slots:
	void OnAddImage();
	void OnSave();
	void OnLoad();

private:
	void CreateToolbar();
	void CreateMenu();

	ViewModel* m_vm;
};