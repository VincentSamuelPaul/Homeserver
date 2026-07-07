// import Iphone15Pro from "../components/magicui/iphone-15-pro"
// import techIcons from "../comp/icons.json"

import { Pointer } from "../components/magicui/pointer"

const Projects = () => {
  return (
    <div className="mt-16 lg:mt-28 w-full jakarta">
      {/* <div className='flex w-full mt-12 items-start justify-start border border-t-white opacity-40'/>   */}
      <h1 className="relative z-20 bg-gradient-to-b text-[#d1d1d1] bg-clip-text py-4 text-4xl lg:text-5xl font-medium text-left">Spotlight</h1>
      <h1 className="lg:mt-4 relative z-20 bg-gradient-to-b text-[#d1d1d1] bg-clip-text py-4 text-md lg:text-xl font-extralight text-left"><span className="font-semibold">Ideas, brought to life.</span> A selection of work that blends clean design with purposeful code—built to work beautifully, and last.</h1>
      {/* <Iphone15Pro 
      className="dark"
        src=""
        /> */}
        <a className="google-sans-code" href="https://ekinshoes.netlify.app/" target="blank">
            <Pointer>
                <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" fill="currentColor" className="bi bi-arrow-up-right-circle hidden lg:block" viewBox="0 0 16 16">
                <path fill-rule="evenodd" d="M1 8a7 7 0 1 0 14 0A7 7 0 0 0 1 8m15 0A8 8 0 1 1 0 8a8 8 0 0 1 16 0M5.854 10.803a.5.5 0 1 1-.708-.707L9.243 6H6.475a.5.5 0 1 1 0-1h3.975a.5.5 0 0 1 .5.5v3.975a.5.5 0 1 1-1 0V6.707z"/>
                </svg>
            </Pointer>
            <div className="mt-12 w-full lg:w-5xl h-40 gradient-1 flex flex-row justify-between clip-b-r project-card p-4">
                <div className="flex flex-col justify-between items-baseline">
                    <h1 className="text-4xl lg:text-5xl font-bold">Ekin Shoes</h1>
                    <h1 className="lg:text-xl text-left">Ecommerce store with full CRUD operations</h1>
                </div>
                <div className="flex flex-col justify-between items-end">
                    <h1 className="text-sm text-right">full stack project</h1>
                </div>
            </div>
        </a>
        <a href="https://github.com/VincentSamuelPaul/gofleet" target="blank">
            <Pointer>
                <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" fill="currentColor" className="bi bi-arrow-up-right-circle hidden lg:block" viewBox="0 0 16 16">
                <path fill-rule="evenodd" d="M1 8a7 7 0 1 0 14 0A7 7 0 0 0 1 8m15 0A8 8 0 1 1 0 8a8 8 0 0 1 16 0M5.854 10.803a.5.5 0 1 1-.708-.707L9.243 6H6.475a.5.5 0 1 1 0-1h3.975a.5.5 0 0 1 .5.5v3.975a.5.5 0 1 1-1 0V6.707z"/>
                </svg>
            </Pointer>
        <div className="flex flex-row justify-between">
            <div/>
            <div className="mt-12 w-4xl h-40 gradient-2 flex flex-row justify-between clip-b-l project-card-r p-4 google-sans-code">
                <div className="flex flex-col justify-between items-end">
                    <h1 className="text-sm lg:text-xl text-left">backend project</h1>
                </div>
                <div className="flex flex-col justify-between items-baseline text-right">
                    <div className="w-full flex flex-row justify-between">
                        <div/>
                        <h1 className="text-4xl lg:text-5xl text-right font-bold">GoFleet</h1>
                    </div>
                    <h1 className="lg:text-xl">Round robin Load balancer</h1>
                </div>
            </div>
        </div>
        </a>
        <a href="https://whirlsearch.netlify.app/" target="blank">
            <Pointer>
                <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" fill="currentColor" className="bi bi-arrow-up-right-circle hidden lg:block" viewBox="0 0 16 16">
                <path fill-rule="evenodd" d="M1 8a7 7 0 1 0 14 0A7 7 0 0 0 1 8m15 0A8 8 0 1 1 0 8a8 8 0 0 1 16 0M5.854 10.803a.5.5 0 1 1-.708-.707L9.243 6H6.475a.5.5 0 1 1 0-1h3.975a.5.5 0 0 1 .5.5v3.975a.5.5 0 1 1-1 0V6.707z"/>
                </svg>
            </Pointer>
        <div className="mt-12 w-full lg:w-5xl h-40 gradient-3 flex flex-row justify-between clip-b-r project-card p-4 google-sans-code">
            <div className="flex flex-col justify-between items-baseline">
                <h1 className="text-4xl lg:text-5xl font-bold">Whirl</h1>
                <h1 className="lg:text-xl text-left">Cricle to Search app for macOS</h1>
            </div>
            <div className="flex flex-col justify-between items-end">
                <h1 className="text-sm lg:text-xl text-right">full stack project</h1>
            </div>
        </div>
        </a>
    </div>
  )
}

export default Projects