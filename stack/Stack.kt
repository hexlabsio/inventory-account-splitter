import io.hexlabs.kloudformation.module.serverless.Method
import io.hexlabs.kloudformation.module.serverless.serverless
import io.kloudformation.KloudFormation
import io.kloudformation.StackBuilder

class Stack: StackBuilder {
    override fun KloudFormation.create(args: List<String>) {
        serverless("account-splitter", "dev", +"hexlabs-deployments") {
            globalRole { roleName("inventory-account-splitter-role") }
            serverlessFunction("account-splitter", +args.first(),+"handler", +"go1.x"){
                http(cors = true) {
                    path("/") { Method.GET(); Method.POST() }
                }
            }
        }
    }
}