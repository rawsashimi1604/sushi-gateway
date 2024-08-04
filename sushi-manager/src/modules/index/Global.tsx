import PluginDropdown from "../../components/sushi-gateway/PluginDropdown";
import Header from "../../components/typography/Header";

function Global() {
    return (
        <section>
            <Header text="Global" align="left" size="md" />

            {/* Gateway metadata */}
            <div className="flex items-center gap-2 text-sm">
                <span>Name</span>
                <span>some-gateway-name</span>
            </div>

            {/* Gateway plugins */}
            <div>
                <span className="text-sm">plugins:</span>
                {/* Some dropdown for plugin design */}
                <PluginDropdown name="basic_auth" data={{}} />
            </div>

        </section>
    );
}

export default Global;
