import { useLocation } from "react-router-dom";
import useBreadcrumbs from "use-react-router-breadcrumbs";

function Breadcrumbs() {
  const location = useLocation();
  const breadcrumbs = useBreadcrumbs();

  return (
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
  );
}

export default Breadcrumbs;
