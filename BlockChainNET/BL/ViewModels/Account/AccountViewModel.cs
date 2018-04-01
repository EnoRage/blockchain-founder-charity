using System;
using System.Windows.Input;
using BlockChainNET.BL.ViewModels;
using BlockChainNET.Helpers;

namespace BlockChainNET.BL.ViewModels.Main
{
    public class AccountViewModel : BaseViewModel
    {
        public ICommand ExitCommand => MakeCommand(() =>
        {
            Settings.GlobalUser = null;
            NavigateTo(Pages.Auth, NavigationMode.RootPage, newNavigationStack: true);
        });

    }
}
