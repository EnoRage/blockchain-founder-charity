using System;
using System.Windows.Input;
using BlockChainNET.BL.ViewModels;

namespace BlockChainNET.BL.ViewModels.Main
{
    public class InfoViewModel : BaseViewModel
    {
        //public ICommand GoToAccountCommand => MakeNavigateToCommand(Pages.Account, NavigationMode.Normal, newNavigationStack: false);
        public ICommand GoToAccountCommand => MakeNavigateToCommand(Pages.Account, NavigationMode.Normal, newNavigationStack: false, withAnimation: true);
    }
}
