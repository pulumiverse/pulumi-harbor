// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumiverse.harbor;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumiverse.harbor.ProjectArgs;
import com.pulumiverse.harbor.Utilities;
import com.pulumiverse.harbor.inputs.ProjectState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.harbor.Project;
 * import com.pulumi.harbor.ProjectArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var main = new Project(&#34;main&#34;, ProjectArgs.builder()        
 *             .enableContentTrust(true)
 *             .public_(false)
 *             .vulnerabilityScanning(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Harbor project example as proxy cache
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.harbor.Registry;
 * import com.pulumi.harbor.RegistryArgs;
 * import com.pulumi.harbor.Project;
 * import com.pulumi.harbor.ProjectArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var docker = new Registry(&#34;docker&#34;, RegistryArgs.builder()        
 *             .providerName(&#34;docker-hub&#34;)
 *             .endpointUrl(&#34;https://hub.docker.com&#34;)
 *             .build());
 * 
 *         var main = new Project(&#34;main&#34;, ProjectArgs.builder()        
 *             .registryId(docker.registryId())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Harbor project can be imported using the `project id` eg, `
 * 
 * ```sh
 *  $ pulumi import harbor:index/project:Project main /projects/1
 * ```
 * 
 *  `
 * 
 */
@ResourceType(type="harbor:index/project:Project")
public class Project extends com.pulumi.resources.CustomResource {
    /**
     * Project allowlist allows vulnerabilities in this list to be ignored in this project when pushing and pulling images. Should be in the format or `[&#34;CVE-123&#34;, &#34;CVE-145&#34;]` or `[&#34;CVE-123&#34;]`
     * 
     */
    @Export(name="cveAllowlists", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> cveAllowlists;

    /**
     * @return Project allowlist allows vulnerabilities in this list to be ignored in this project when pushing and pulling images. Should be in the format or `[&#34;CVE-123&#34;, &#34;CVE-145&#34;]` or `[&#34;CVE-123&#34;]`
     * 
     */
    public Output<Optional<List<String>>> cveAllowlists() {
        return Codegen.optional(this.cveAllowlists);
    }
    /**
     * Prevent deployment of images with vulnerability severity equal or higher than the specified value. Images must be scanned before this takes effect. Possible values: `critical`, `high`, `medium`, `low`, `none`. (Default: `&#34;&#34;` - empty)
     * 
     */
    @Export(name="deploymentSecurity", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deploymentSecurity;

    /**
     * @return Prevent deployment of images with vulnerability severity equal or higher than the specified value. Images must be scanned before this takes effect. Possible values: `critical`, `high`, `medium`, `low`, `none`. (Default: `&#34;&#34;` - empty)
     * 
     */
    public Output<Optional<String>> deploymentSecurity() {
        return Codegen.optional(this.deploymentSecurity);
    }
    /**
     * Enables Content Trust for project. When enabled it queries the embedded docker notary server. Can be set to `&#34;true&#34;` or `&#34;false&#34;` (Default: false)
     * 
     */
    @Export(name="enableContentTrust", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableContentTrust;

    /**
     * @return Enables Content Trust for project. When enabled it queries the embedded docker notary server. Can be set to `&#34;true&#34;` or `&#34;false&#34;` (Default: false)
     * 
     */
    public Output<Optional<Boolean>> enableContentTrust() {
        return Codegen.optional(this.enableContentTrust);
    }
    /**
     * A boolean that indicates all repositories should be deleted from the project so that the project can be destroyed without error. These repositories are *not* recoverable.
     * 
     */
    @Export(name="forceDestroy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> forceDestroy;

    /**
     * @return A boolean that indicates all repositories should be deleted from the project so that the project can be destroyed without error. These repositories are *not* recoverable.
     * 
     */
    public Output<Optional<Boolean>> forceDestroy() {
        return Codegen.optional(this.forceDestroy);
    }
    /**
     * The name of the project that will be created in harbor.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the project that will be created in harbor.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The id of the project with harbor.
     * 
     */
    @Export(name="projectId", refs={Integer.class}, tree="[0]")
    private Output<Integer> projectId;

    /**
     * @return The id of the project with harbor.
     * 
     */
    public Output<Integer> projectId() {
        return this.projectId;
    }
    /**
     * The project will be public accessibility. Can be set to `&#34;true&#34;` or `&#34;false&#34;` (Default: false)
     * 
     */
    @Export(name="public", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> public_;

    /**
     * @return The project will be public accessibility. Can be set to `&#34;true&#34;` or `&#34;false&#34;` (Default: false)
     * 
     */
    public Output<Optional<String>> public_() {
        return Codegen.optional(this.public_);
    }
    /**
     * To enabled project as Proxy Cache
     * 
     */
    @Export(name="registryId", refs={Integer.class}, tree="[0]")
    private Output<Integer> registryId;

    /**
     * @return To enabled project as Proxy Cache
     * 
     */
    public Output<Integer> registryId() {
        return this.registryId;
    }
    /**
     * The storage quota of the project in GB&#39;s
     * 
     */
    @Export(name="storageQuota", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> storageQuota;

    /**
     * @return The storage quota of the project in GB&#39;s
     * 
     */
    public Output<Optional<Integer>> storageQuota() {
        return Codegen.optional(this.storageQuota);
    }
    /**
     * Images will be scanned for vulnerabilities when push to harbor. Can be set to `&#34;true&#34;` or `&#34;false&#34;` (Default: true)
     * 
     */
    @Export(name="vulnerabilityScanning", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> vulnerabilityScanning;

    /**
     * @return Images will be scanned for vulnerabilities when push to harbor. Can be set to `&#34;true&#34;` or `&#34;false&#34;` (Default: true)
     * 
     */
    public Output<Optional<Boolean>> vulnerabilityScanning() {
        return Codegen.optional(this.vulnerabilityScanning);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Project(String name) {
        this(name, ProjectArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Project(String name, @Nullable ProjectArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Project(String name, @Nullable ProjectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("harbor:index/project:Project", name, args == null ? ProjectArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Project(String name, Output<String> id, @Nullable ProjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("harbor:index/project:Project", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Project get(String name, Output<String> id, @Nullable ProjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Project(name, id, state, options);
    }
}
