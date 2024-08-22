import { BsGithub } from "react-icons/bs";
import { Link } from "react-router-dom";
import { useLocation } from "react-router-dom";
import useBreadcrumbs from "use-react-router-breadcrumbs";

function Breadcrumbs() {
  const location = useLocation();
  const breadcrumbs = useBreadcrumbs();

  return (
    <div className="flex items-center justify-between w-full">
      <div className="flex flex-row items-center text-sm gap-1">
        <span className="text-gray-500">gateway {">"}</span>
        {breadcrumbs.map((breadcrumb, i) => {
          return (
            <span
              key={i + (breadcrumb.breadcrumb?.toString() as string)}
              className={`${
                breadcrumb.key === location.pathname
                  ? "font-semibold"
                  : "text-gray-500"
              }`}
            >
              {breadcrumb.key === "/" ? "manager" : breadcrumb.key.slice(1)}{" "}
              {i + 1 !== breadcrumbs.length && ">"}
            </span>
          );
        })}
      </div>
      <Link
        to="https://github.com/rawsashimi1604/sushi-gateway"
        target="_blank"
      >
        <BsGithub className="cursor-pointer duration-150 transition-all hover:opacity-80" />
      </Link>
    </div>
  );
}

export default Breadcrumbs;
