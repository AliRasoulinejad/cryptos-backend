package app

import (
	"fmt"

	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

// AppName is the app's name .
var (
	// GitCommit is SHA1 ref of current build
	GitCommit string

	// GitRef is branch name of current build
	GitRef string

	// GitTag is version name of current build
	GitTag string

	// BuildDate is the timestamp of build
	BuildDate string

	// CompilerVersion is the version of go compiler
	CompilerVersion string
)

const AppName string = "cryptos_backend"

const asciiArt = `
        CCCCCCCCCCCCC                                                                       tttt                                            
     CCC::::::::::::C                                                                    ttt:::t                                            
   CC:::::::::::::::C                                                                    t:::::t                                            
  C:::::CCCCCCCC::::C                                                                    t:::::t                                            
 C:::::C       CCCCCCrrrrr   rrrrrrrrryyyyyyy           yyyyyyyppppp   ppppppppp   ttttttt:::::ttttttt       ooooooooooo       ssssssssss   
C:::::C              r::::rrr:::::::::ry:::::y         y:::::y p::::ppp:::::::::p  t:::::::::::::::::t     oo:::::::::::oo   ss::::::::::s  
C:::::C              r:::::::::::::::::ry:::::y       y:::::y  p:::::::::::::::::p t:::::::::::::::::t    o:::::::::::::::oss:::::::::::::s 
C:::::C              rr::::::rrrrr::::::ry:::::y     y:::::y   pp::::::ppppp::::::ptttttt:::::::tttttt    o:::::ooooo:::::os::::::ssss:::::s
C:::::C               r:::::r     r:::::r y:::::y   y:::::y     p:::::p     p:::::p      t:::::t          o::::o     o::::o s:::::s  ssssss 
C:::::C               r:::::r     rrrrrrr  y:::::y y:::::y      p:::::p     p:::::p      t:::::t          o::::o     o::::o   s::::::s      
C:::::C               r:::::r               y:::::y:::::y       p:::::p     p:::::p      t:::::t          o::::o     o::::o      s::::::s   
 C:::::C       CCCCCC r:::::r                y:::::::::y        p:::::p    p::::::p      t:::::t    tttttto::::o     o::::ossssss   s:::::s 
  C:::::CCCCCCCC::::C r:::::r                 y:::::::y         p:::::ppppp:::::::p      t::::::tttt:::::to:::::ooooo:::::os:::::ssss::::::s
   CC:::::::::::::::C r:::::r                  y:::::y          p::::::::::::::::p       tt::::::::::::::to:::::::::::::::os::::::::::::::s 
     CCC::::::::::::C r:::::r                 y:::::y           p::::::::::::::pp          tt:::::::::::tt oo:::::::::::oo  s:::::::::::ss  
        CCCCCCCCCCCCC rrrrrrr                y:::::y            p::::::pppppppp              ttttttttttt     ooooooooooo     sssssssssss    
                                            y:::::y             p:::::p                                                                     
                                           y:::::y              p:::::p                                                                     
                                          y:::::y              p:::::::p                                                                    
                                         y:::::y               p:::::::p                                                                    
                                        yyyyyyy                p:::::::p                                                                    
                                                               ppppppppp
`

type Application struct {
	DB           *gorm.DB
	Repositories *Repositories
	Tracer       trace.Tracer
}

func Banner() string {
	return asciiArt + "\n" + fmt.Sprintf("Tag: %s, Ref: %s, GitCommit: %s, BuildDate: %s, Compiler: %s", GitTag, GitRef, GitCommit, BuildDate, CompilerVersion)
}
