import PluginDropdown from "../../components/sushi-gateway/PluginDropdown";
import Header from "../../components/typography/Header";
import Tag from "../../components/typography/Tag";
import NormalText from "../../components/typography/Text";

function Global() {
    return (
        <section className="">
            <Header text="Global" align="left" size="md" />

            {/* Gateway metadata */}
            <div className="bg-neutral-200 px-4 py-2 rounded-lg shadow-sm w-[80%]">
                <div className="flex items-center gap-2 text-sm">
                    <Tag value="name" />
                    <NormalText text="some-gateway-name" />
                </div>

                {/* Gateway plugins */}
                <div>
                    <div className="mb-2">
                        <Tag value="plugins" />
                    </div>

                    {/* Some dropdown for plugin design */}
                    <PluginDropdown name="http_log" data={{
                        "name": "http_log",
                        "enabled": true,
                        "data": {
                            "http_endpoint": "http://localhost:8003/v1/log",
                            "method": "POST",
                            "content_type": "application/json"
                        }
                    }} />
                </div>
            </div>


        </section>
    );
}

export default Global;
