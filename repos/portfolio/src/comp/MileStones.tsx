import { Pointer } from "../components/magicui/pointer"

const MileStones = () => {
  return (
    <div className="mt-16 lg:mt-32 jakarta">
        <h1 className=" relative z-20 bg-gradient-to-b text-[#d1d1d1] bg-clip-text py-4 text-4xl lg:text-5xl font-medium text-left ">Milestones</h1>
        <a className="mt-4 border-b-2 border-[#717171] pb-4" href="https://www.linkedin.com/posts/thatsvincentpaul_hackathon-innovation-aiforgood-activity-7328847771988295681-_1FH?utm_source=share&utm_medium=member_desktop&rcm=ACoAADaYWOcBS6RpoP2NDLvxsszyesH8jxK0Uqk" target="blank">
            <Pointer>
                <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" fill="white" className="bi bi-arrow-up-right-circle hidden lg:block" viewBox="0 0 16 16">
                <path fill-rule="evenodd" d="M1 8a7 7 0 1 0 14 0A7 7 0 0 0 1 8m15 0A8 8 0 1 1 0 8a8 8 0 0 1 16 0M5.854 10.803a.5.5 0 1 1-.708-.707L9.243 6H6.475a.5.5 0 1 1 0-1h3.975a.5.5 0 0 1 .5.5v3.975a.5.5 0 1 1-1 0V6.707z"/>
                </svg>
            </Pointer>
            <div className="flex flex-col lg:flex-row justify-between items-start">
                <h1 className="relative z-20 bg-gradient-to-b text-[#d1d1d1] bg-clip-text py-4 text-3xl lg:text-4xl font-extralight text-left">Quantum Breach <span className=" italic">- winner</span></h1>
                <div className="flex flex-col justify-baseline items-start">
                    <a href="https://www.cloudsek.com/" target="blank" className="w-full relative z-20 bg-gradient-to-b text-[#d1d1d1] bg-clip-text text-xl lg:text-xl font-extralight text-left lg:text-right hover:cursor-pointer hover:underline">organised by CloudSEK ↗</a>
                    <h1 className="relative z-20 bg-gradient-to-b text-[#d1d1d1] bg-clip-text text-lg lg:text-lg font-extralight text-left lg:text-right ">May 2025 - AI + Cybersecurity</h1>
                </div>
            </div>
            <h1 className="relative z-20 bg-gradient-to-b text-[#d1d1d1] bg-clip-text py-4 text-md lg:text-xl font-extralight text-left"><span className="font-semibold">Key Guardian</span> - Intelligence that protects. Every keystroke is analyzed in real time to detect harmful content, with instant, live updates that keep parents informed—anywhere, anytime.</h1>
        </a>
        <a className="mt-12 border-b-2 border-[#717171] pb-4" href="https://www.linkedin.com/posts/thatsvincentpaul_ai-cybersecurity-hackathonwinner-activity-7322135813528391680-5JvM?utm_source=share&utm_medium=member_desktop&rcm=ACoAADaYWOcBS6RpoP2NDLvxsszyesH8jxK0Uqk" target="blank">
            <Pointer>
                <svg xmlns="http://www.w3.org/2000/svg" width="40" height="40" fill="white" className="bi bi-arrow-up-right-circle hidden lg:block" viewBox="0 0 16 16">
                <path fill-rule="evenodd" d="M1 8a7 7 0 1 0 14 0A7 7 0 0 0 1 8m15 0A8 8 0 1 1 0 8a8 8 0 0 1 16 0M5.854 10.803a.5.5 0 1 1-.708-.707L9.243 6H6.475a.5.5 0 1 1 0-1h3.975a.5.5 0 0 1 .5.5v3.975a.5.5 0 1 1-1 0V6.707z"/>
                </svg>
            </Pointer>
            <div className="flex flex-col lg:flex-row justify-between items-start">
                <h1 className="relative z-20 bg-gradient-to-b text-[#d1d1d1] bg-clip-text py-4 text-3xl lg:text-4xl font-extralight text-left">Xenith <span className=" italic">- runnerup</span></h1>
                <div className="flex flex-col justify-baseline items-start">
                    <a href="https://dsatm.edu.in/" target="blank" className="w-full relative z-20 bg-gradient-to-b text-[#d1d1d1] bg-clip-text text-xl lg:text-xl font-extralight text-left lg:text-right hover:cursor-pointer hover:underline">organised by DSATM ↗</a>
                    <h1 className="relative z-20 bg-gradient-to-b text-[#d1d1d1] bg-clip-text text-lg lg:text-lg font-extralight text-left lg:text-right ">June 2025 - Remote Area Solutions</h1>
                </div>
            </div>
            <h1 className="relative z-20 bg-gradient-to-b text-[#d1d1d1] bg-clip-text py-4 text-md lg:text-xl font-extralight text-left"><span className="font-semibold">VoiceBridge</span> - an AI-powered voice assistant designed to help people in remote areas—especially farmers—get answers to their questions using only a basic phone call. No smartphones. No internet. Just a phone call.</h1>
        </a>
    </div>
  )
}

export default MileStones