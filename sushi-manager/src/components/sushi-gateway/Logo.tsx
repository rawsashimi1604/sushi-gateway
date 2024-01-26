import { ReactComponent as LogoSvg } from "../../assets/sushi.svg";

function Logo() {
  return (
    <div className="flex flex-row justify-center items-center font-light text-2xl font-sans gap-4">
      <LogoSvg className="w-16 h-16" />
      <div className="relative">
        <h1 className="-mt-2 font-lora font-ligh tracking-widest">
          SUSHI GATEWAY
        </h1>
        <div className="absolute right-0 top-5.5">
          <span className="font-lora tracking-wider text-sm border border-black rounded-sm shadow-md px-2 py-0.5 pb-1 font-bold">
            manager
          </span>
        </div>
      </div>
    </div>
  );
}

export default Logo;
