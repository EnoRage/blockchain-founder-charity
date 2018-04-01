namespace BlockChainNET
{
    public enum Pages
    {
        Account,
        Auth,
        Info
    }

    public enum NavigationMode
    {
        Normal,
        Modal,
        RootPage,
        Custom
    }

    public enum PageState
    {
        Clean,
        Loading,
        Normal,
        NoData,
        Error,
        NoInternet
    }
}
