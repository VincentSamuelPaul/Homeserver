import techStack from "./icons.json";

const TechStack = () => {
  return (
    <div className="google-sans-code">
      <div className='flex w-full mt-12 items-start justify-start border border-t-white opacity-40'/>  
      {/* <h1 className="mt-12 relative z-20 bg-gradient-to-b text-white bg-clip-text py-4 text-3xl lg:text-5xl font-light text-left">powered by.</h1>  */}
      <div className="">
        <div className="mt-12 flex flex-row justify-start items-center">
          <div className="w-full grid grid-cols-[repeat(auto-fit,minmax(3.5rem,1fr))] gap-6">
            {techStack.flatMap(category => category.items).map((stack, idx) => (
              <div
                key={idx}
                className="flex items-center justify-center w-14 h-14 lg:w-20 lg:h-20 transition-all duration-500 transform grayscale hover:grayscale-0 hover:scale-110"
                dangerouslySetInnerHTML={{ __html: stack.svg }}
              />
            ))}
          </div>
        </div>
      </div>
    </div> 
  )
}

export default TechStack