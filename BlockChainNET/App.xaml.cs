using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using BlockChainNET.UI;
using BlockChainNET;
using Xamarin.Forms;
using Xamarin.Forms.Xaml;

[assembly: XamlCompilation(XamlCompilationOptions.Compile)]

namespace BlockChainNET
{
    public partial class App : Application
    {
        public App()
        {
            InitializeComponent();

            NavigationService.Init(this);
            DialogService.Init(this);
            NavigationService.Instance.SetMainPage(Pages.Auth);
            //DataServices.Init(true);
            /*if(Settings.GetPacientSetting() == null || string.IsNullOrWhiteSpace(Settings.GetPacientSetting().Id))
            {
                NavigationService.Instance.SetMainPage(Pages.Auth);
                //NavigationService.Instance.Push(Pages.Auth);
            }
            else
            {
                NavigationService.Instance.SetMainMasterDetailPage(Pages.Menu, Pages.Info);
            }*/
        }

        protected override void OnStart()
        {
            // Handle when your app starts
        }

        protected override void OnSleep()
        {
            // Handle when your app sleeps
        }

        protected override void OnResume()
        {
            // Handle when your app resumes
        }
    }
}