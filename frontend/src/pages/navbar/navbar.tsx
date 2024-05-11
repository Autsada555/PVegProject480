
import iconprofile from '../../assets/iconprofile.png'

const Navbar = () => {

    return (
        <>
            <div className="relative">
                <div className="w-full h-[90px] bg-[#01BD63] flex items-center justify-between px-4 sm:px-6 lg:px-8">
                    <div className="text-white text-xl font-normal font-['Inter'] leading-7">Food Logo</div>

                    <nav className="flex space-x-12 ">
                        <div>
                            <input
                                type="text"
                                placeholder="search..."
                                className="px-4 py-2 rounded border border-gray-300 focus:outline-none focus:border-blue-500"
                            />
                        </div>
                        <a href="#" className="text-white text-xl font-normal font-['Inter'] leading-9">Home</a>
                        <a href="#" className="text-white text-xl font-normal font-['Inter'] leading-9">Menu</a>
                        <a href="#" className="text-white text-xl font-normal font-['Inter'] leading-9">About</a>
                        <a href="#" className="text-white text-xl font-normal font-['Inter'] leading-9">Contract</a>
                        <div>
                            <img className="absolute right-0 top-3 w-12 h-12 mr-4 mt-2" src={iconprofile} alt="icon" />
                        </div>
                    </nav>

                </div>

            </div>
        </>
    );
};

export default Navbar;
