import { ReactComponent as LogoSvg } from "../../assets/sushi.svg";
import logo from "../../assets/logos/sushi-manager.png";

function Logo() {
  return (
    <div className="w-60">
      <img src={logo} />
    </div>
  );
}

export default Logo;
