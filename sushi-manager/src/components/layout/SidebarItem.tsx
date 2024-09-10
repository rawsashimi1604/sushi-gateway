import { Link, useLocation } from "react-router-dom";

interface SidebarItemProps {
  item: string;
  href?: string;
}

function SidebarItem({ item, href = "/" }: SidebarItemProps) {
  const location = useLocation();
  const isSelected = location.pathname == href;

  return (
    <li
      className={`${
        isSelected && "bg-blue-100 py-0.5 pl-2 font-semibold"
      } hover:bg-slate-100 transition-all hover:pl-4 font-lora duration-150 hover:cursor-pointer rounded-lg tracking-widest`}
    >
      <Link to={href}>
        <span className="hover:transition-all duration-150">{item}</span>
      </Link>
    </li>
  );
}

export default SidebarItem;
