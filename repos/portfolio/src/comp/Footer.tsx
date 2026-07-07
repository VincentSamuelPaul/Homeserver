import { Pointer } from "../components/magicui/pointer"

const Footer = () => {
  return (
    <div className="mt-16 lg:mt-28 text-white jakarta">
      <Pointer className='fill-[#c4c4c4]'/>
        <h1 className="relative z-20 bg-gradient-to-b text-[#d1d1d1] bg-clip-text py-4 text-3xl lg:text-5xl font-medium text-left">Wrapping up</h1>
        {/* <h1 className="relative z-20 bg-gradient-to-b text-[#d1d1d1] bg-clip-text py-4 text-md lg:text-xl font-medium text-left">Product-focused developer building reliable, scalable backend systems — in Go. Passionate about solving real problems, crafting clean APIs, and exploring distributed systems. I also design high-converting landing pages with React and Tailwind.</h1> */}
        <h1 className="lg:mt-4 relative z-20 bg-gradient-to-b text-[#d1d1d1] bg-clip-text py-4 text-2xl lg:text-3xl font-extralight text-left italic opacity-90">Quote ~</h1>
        <h1 className="text-left opacity-50">❝</h1>
        <h1 className="py-4 relative z-20 bg-gradient-to-b text-[#d1d1d1] bg-clip-text text-xl lg:text-3xl font-extralight text-left italic opacity-50">It's not about the language or the framework, magic is in my fingertips.</h1>
        <h1 className="text-left opacity-50">❞</h1>
        <h1 className="lg:mt-4 relative z-20 bg-gradient-to-b text-[#d1d1d1] bg-clip-text py-4 text-xl lg:text-2xl font-extralight text-left">Produced by - <span className="font-semibold italic">Vincent Paul</span></h1>
    </div>
  )
}

export default Footer