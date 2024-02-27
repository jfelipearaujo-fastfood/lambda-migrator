from diagrams import Cluster, Diagram
from diagrams.aws.database import RDSPostgresqlInstance as RDS
from diagrams.onprem.ci import GithubActions
from diagrams.aws.compute import LambdaFunction as Lambda

diagram_attr = {
    "fontsize": "25",
    "bgcolor": "white",
    "pad": "0.5",
    "splines": "spline",
}

cluster_attr = {
    "fontsize": "15",
    "size": "5",
    "margin": "20",
    "pad": "2"
}

item_attr = {
    "fontsize": "15",
    "height": "2.2"
}

diagram_label = """

Cloud AWS Database Migrations

"""

with Diagram(diagram_label, show=False, graph_attr=diagram_attr, direction="LR"):
    github = GithubActions("Github Actions", **item_attr)

    with Cluster("AWS", graph_attr=cluster_attr):
        lambda_function = Lambda("Lambda Migrator", **item_attr)
        rds = RDS("RDS", **item_attr)

        github >> lambda_function
        
        lambda_function >> rds