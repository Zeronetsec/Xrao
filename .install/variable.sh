export prefix="${PREFIX:-/usr}"; readonly prefix

export tmp="${prefix}/tmp"; readonly tmp
export opt="${prefix}/opt"; readonly opt
export bin="${prefix}/bin"; readonly bin

export bkdate="$(
    command date '+%Y_%b_%d_%H_%M_%S'
)"; readonly bkdate