using System;
using System.Windows.Input;
using BlockChainNET.BL.DB;
using BlockChainNET.BL.ViewModels;
using BlockChainNET.Helpers;

namespace BlockChainNET.BL.ViewModels.Main
{
    public class AuthViewModel : BaseViewModel
    {
        public ICommand GoToInfoCommand => MakeCommand(async () =>
        {
            ShowLoading();
            var resp = await DataBaseService.UserAuthCheck("@maxspT");
            if (resp)
            {
                Settings.GlobalUser = await DataBaseService.GetUser("@maxspT");
            }
            else
            {
                await ShowAlert("Warning", "User Not Found", "Ok");
            }
            HideLoading();

            if (resp)
            {
                NavigateTo(Pages.Info, NavigationMode.RootPage, newNavigationStack: true, withAnimation: true);
            }
        });
    }
}
