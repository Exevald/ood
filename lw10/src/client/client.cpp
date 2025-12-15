#include <QApplication>
#include "../viewModel/ViewModel.h"
#include "../view/MainWindow.h"

int RunClient(int argc, char** argv)
{
	QApplication app(argc, argv);
	const auto vm = new ViewModel();

	MainWindow window(vm);
	window.show();

	return QApplication::exec();
}