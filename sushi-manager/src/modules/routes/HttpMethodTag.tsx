import { HttpMethod } from "../../types/Http";
import Tag from "../../components/typography/Tag";

export interface HttpMethodTag {
  method: HttpMethod;
}

// TODO: fix colors...
function HttpMethodTag({ method }: HttpMethodTag) {
  const colorMap: { [key: string]: string } = {
    GET: "bg-httpGet",
    POST: "bg-httpPost",
    PUT: "bg-httpPut",
    PATCH: "bg-httpPatch",
    DELETE: "bg-httpDelete",
    OPTIONS: "bg-httpOptions",
  };

  return (
    <Tag
      value={method}
      bgColor={colorMap[method]}
      className="text-white shadow-md"
    />
  );
}

export default HttpMethodTag;
