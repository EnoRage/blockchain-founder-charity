using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using BlockChainMobileBack.Models;
using Microsoft.AspNetCore.Mvc;
using BlockChainMobileBack.Repository;

// For more information on enabling Web API for empty projects, visit https://go.microsoft.com/fwlink/?LinkID=397860

namespace BlockChainMobileBack.Controllers
{
    [Route("api/[controller]")]
    public class HelpController : Controller
    {
        private readonly IDataBaseRepository _dataBaseRepository;

        public HelpController(IDataBaseRepository dataBaseRepository)
        {
            _dataBaseRepository = dataBaseRepository;
        }

        [HttpGet("{setting}")]
        public async Task<string> Get(string setting)
        {
            if (setting == "init")
            {
                FoundationOptions foundationOptions = new FoundationOptions()
                {
                    Capital = 0.0,
                    FoundedDate = DateTime.Now,
                    Mission = "Ne Otsosat",
                    Name = "MIREA",
                    Country = "Ukraine"
                };
                await _dataBaseRepository.AddFoundationOption(foundationOptions);
            }
            return "Done";
        }
    }
}
